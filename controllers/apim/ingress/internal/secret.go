// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gravitee-io/gravitee-kubernetes-operator/internal/indexer"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/gravitee-io/gravitee-kubernetes-operator/pkg/keys"
	core "k8s.io/api/core/v1"
	netV1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	util "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (d *Delegate) updateIngressTLSReference(ingress *netV1.Ingress) error {
	if ingress.Spec.TLS == nil || len(ingress.Spec.TLS) == 0 {
		d.log.Info("no TLS will be configured")
		return nil
	}

	/*
		// We can also check if the TLS is already has a reference inside the rules.host section of the ingress
		for _, rule := range ingress.Spec.Rules
			for _, tls := range ingress.Spec.TLS
				for _, host := range tls.Hosts
					if rule.Host == host
	*/
	key := fmt.Sprintf("%s-%s", ingress.Namespace, ingress.Name)
	values := make([]string, 0)
	for _, tls := range ingress.Spec.TLS {
		secret := &core.Secret{}
		key := types.NamespacedName{Namespace: ingress.Namespace, Name: tls.SecretName}
		if err := d.k8s.Get(d.ctx, key, secret); err != nil {
			return err
		}

		// If finalizer not present, add it; This is a new object
		if !util.ContainsFinalizer(secret, keys.KeyPairFinalizer) {
			d.log.Info("adding finalizer to the tls secret")

			secret.ObjectMeta.Finalizers = append(secret.ObjectMeta.Finalizers, keys.KeyPairFinalizer)
			if err := d.k8s.Update(d.ctx, secret); err != nil {
				return client.IgnoreNotFound(err)
			}
		}

		// Secret has been deleted while it still has reference to an ingress, not allowed
		if !secret.DeletionTimestamp.IsZero() {
			return fmt.Errorf("secret can't be deleted because it has reference to an existing ingress [%s]", ingress.Name)
		}

		// parse the secrete just to make sure the data is valid before
		// passing it to the gateway
		if err := d.parseTLSSecret(secret); err != nil {
			return err
		}

		values = append(values, fmt.Sprintf("%s/%s", secret.Namespace, secret.Name))
	}

	d.log.Info("Update GW PEM registry with the secret names")
	return d.updatePemRegistry(ingress, key, values)
}

func (d *Delegate) deleteIngressTLSReference(ingress *netV1.Ingress) error {
	for _, tls := range ingress.Spec.TLS {
		secret := &core.Secret{}
		key := types.NamespacedName{Namespace: ingress.Namespace, Name: tls.SecretName}
		if err := d.k8s.Get(d.ctx, key, secret); err != nil {
			return err
		}

		// It is possible that the same secret has been used in another ingress
		// We will not remove the finalizer and we will not remove the keypair
		// from keystore but we also don't throw any error to let the current ingress
		// be deleted
		hasReferenceToOtherIngress, err := d.secretHasReference(ingress, secret)
		if err != nil {
			return err
		}

		if hasReferenceToOtherIngress {
			d.log.Error(
				errors.New("secret has reference"),
				"secret is used by another ingress, it will not be deleted from the keystore")
		} else {
			d.log.Info("removing finalizer from secret", "secret", secret.Name)
			util.RemoveFinalizer(secret, keys.KeyPairFinalizer)

			if err = d.k8s.Update(d.ctx, secret); err != nil {
				return err
			}
		}
	}

	// no reference to this secret, we can remove it from the keystore
	key := fmt.Sprintf("%s-%s", ingress.Namespace, ingress.Name)
	if err := d.updatePemRegistry(ingress, key, nil); err != nil {
		return err
	}

	d.log.Info("gateway pem registry has been successfully updated.")
	return nil
}

func (d *Delegate) secretHasReference(ing *netV1.Ingress, secret *core.Secret) (bool, error) {
	il, err := d.retrieveIngressListWithTLS(d.ctx, ing.Namespace)
	if err != nil {
		return false, err
	}

	for i := range il.Items {
		for _, tls := range il.Items[i].Spec.TLS {
			if tls.SecretName == secret.Name && il.Items[i].DeletionTimestamp.IsZero() {
				d.log.Info("the secret is already used inside an ingress resource", "resource", il.Items[i].Name)
				return true, nil
			}
		}
	}

	return false, nil
}

func (d *Delegate) retrieveIngressListWithTLS(ctx context.Context, ns string) (*netV1.IngressList, error) {
	il := &netV1.IngressList{}
	if err := d.k8s.List(ctx, il, client.InNamespace(ns)); err != nil {
		return nil, client.IgnoreNotFound(err)
	}

	result := &netV1.IngressList{}
	for i := range il.Items {
		ingress := il.Items[i]
		if indexer.IsGraviteeIngress(&ingress) {
			if ingress.Spec.TLS != nil {
				result.Items = append(result.Items, ingress)
			}
		}
	}

	return result, nil
}

func (d *Delegate) updatePemRegistry(ing *netV1.Ingress, key string, values []string) error {
	pemRegistriesToUpdate, err := d.getPemRegistryConfigMapsToUpdate(ing)
	if err != nil {
		return err
	}

	if !ing.DeletionTimestamp.IsZero() {
		return d.deletePemRegistryEntry(pemRegistriesToUpdate, key)
	}

	return d.updatePemRegistryEntry(pemRegistriesToUpdate, key, values)
}

func (d *Delegate) getPemRegistryConfigMapsToUpdate(ing *netV1.Ingress) ([]*core.ConfigMap, error) {
	pemRegistryConfigMaps := &core.ConfigMapList{}
	filter := &client.ListOptions{
		LabelSelector: labels.SelectorFromSet(labels.Set{keys.GraviteeComponentLabel: keys.GraviteePemRegistryLabel}),
	}
	var err error
	if err = d.k8s.List(d.ctx, pemRegistryConfigMaps, filter); err != nil {
		return nil, err
	}

	if len(pemRegistryConfigMaps.Items) == 0 {
		return nil, fmt.Errorf("unable to find any pem-registry configmap in the cluster")
	}

	pemRegistriesToUpdate := make([]*core.ConfigMap, 0)
	for i := range pemRegistryConfigMaps.Items {
		item := pemRegistryConfigMaps.Items[i]
		if (ing.Spec.IngressClassName != nil && item.Labels[keys.IngressClassAnnotation] == *ing.Spec.IngressClassName) ||
			item.Labels[keys.IngressClassAnnotation] == ing.GetAnnotations()[keys.IngressClassAnnotation] {
			pemRegistriesToUpdate = append(pemRegistriesToUpdate, &item)
		}
	}

	return pemRegistriesToUpdate, nil
}

// parse K8S TLS secret and make sure it is valid.
func (d *Delegate) parseTLSSecret(secret *core.Secret) error {
	// get the key and certificate (The TLS secret must contain keys named tls.crt and tls.key
	// https://kubernetes.io/docs/concepts/services-networking/ingress/#tls
	pemKeyBytes, ok := secret.Data["tls.key"]
	if !ok {
		return fmt.Errorf("%s", "tls key not found in secret")
	}

	tlsKey, _ := pem.Decode(pemKeyBytes)
	if tlsKey == nil {
		return fmt.Errorf("%s", "can not decode the tls key")
	}

	if !strings.Contains(tlsKey.Type, "PRIVATE KEY") {
		return fmt.Errorf("%s", "wrong tls key type")
	}

	pemCrtBytes, ok := secret.Data["tls.crt"]
	if !ok {
		return fmt.Errorf("%s", "tls cert not found in secret")
	}

	tlsCrt, _ := pem.Decode(pemCrtBytes)
	if tlsCrt == nil {
		return fmt.Errorf("%s", "can not decode the tls certificate")
	}

	_, err := x509.ParseCertificate(tlsCrt.Bytes)
	if err != nil {
		return err
	}

	if tlsCrt.Type != "CERTIFICATE" {
		return fmt.Errorf("%s", "wrong tls certification type")
	}

	return nil
}

func (d *Delegate) updatePemRegistryEntry(configmaps []*core.ConfigMap, key string, values []string) error {
	for _, configmap := range configmaps {
		// a simple solution for dealing with secretes there were updated
		// the gateway will receive an update event and will refresh the trust store
		updateTimestamp(configmap)

		if configmap.Data == nil {
			configmap.Data = map[string]string{}
		}

		bytes, err := json.Marshal(values)
		if err != nil {
			return err
		}

		configmap.Data[key] = string(bytes)

		err = d.k8s.Update(d.ctx, configmap)

		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Delegate) deletePemRegistryEntry(configmaps []*core.ConfigMap, key string) error {
	for _, configmap := range configmaps {
		// a simple solution for dealing with secretes there were updated
		// the gateway will receive an update event and will refresh the trust store
		updateTimestamp(configmap)

		delete(configmap.Data, key)
		if err := d.k8s.Update(d.ctx, configmap); err != nil {
			return err
		}
	}

	return nil
}

func updateTimestamp(pemRegistryConfigMap *core.ConfigMap) {
	pemRegistryConfigMap.Annotations["updateTimestamp"] = time.Now().Format("2006-01-02T15:04:05Z")
}
