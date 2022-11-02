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

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"
	k8sUtil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	gio "github.com/gravitee-io/gravitee-kubernetes-operator/api/v1alpha1"
	"github.com/gravitee-io/gravitee-kubernetes-operator/pkg/keys"
	"github.com/gravitee-io/gravitee-kubernetes-operator/test/internal"
)

var _ = Describe("Checking NoneRecoverable && Recoverable error", Label("DisableSmokeExpect"), func() {

	Context("With basic ApiDefinition & ManagementContext", func() {
		var managementContextFixture *gio.ManagementContext
		var apiDefinitionFixture *gio.ApiDefinition

		var savedApiDefinition *gio.ApiDefinition

		var apiLookupKey types.NamespacedName

		BeforeEach(func() {
			By("Create a management context to synchronize with the REST API")

			apiWithContext, err := internal.NewApiWithRandomContext(
				internal.BasicApiFile, internal.ContextWithSecretFile,
			)

			Expect(err).ToNot(HaveOccurred())

			managementContext := apiWithContext.Context
			Expect(k8sClient.Create(ctx, managementContext)).Should(Succeed())

			By("Create an API definition resource stared by default")

			apiDefinition := apiWithContext.Api
			Expect(k8sClient.Create(ctx, apiDefinition)).Should(Succeed())

			apiDefinitionFixture = apiDefinition
			managementContextFixture = managementContext
			apiLookupKey = types.NamespacedName{Name: apiDefinitionFixture.Name, Namespace: namespace}

			By("Expect the API Definition is Ready")
			savedApiDefinition = new(gio.ApiDefinition)
			Eventually(func() bool {
				k8sErr := k8sClient.Get(ctx, apiLookupKey, savedApiDefinition)
				return k8sErr == nil && savedApiDefinition.Status.CrossID != ""
			}, timeout, interval).Should(BeTrue())
		})

		It("Should not requeue reconcile with 401 error", func() {

			By("Set bad credentials in ManagementContext")
			managementContextBad := managementContextFixture.DeepCopy()
			managementContextBad.Spec.Auth.SecretRef = nil
			managementContextBad.Spec.Auth.BearerToken = "bad-token"

			err := k8sClient.Update(ctx, managementContextBad)
			Expect(err).ToNot(HaveOccurred())

			By("Update the API definition")
			apiDefinition := savedApiDefinition.DeepCopy()
			apiDefinition.Spec.Name = "new-name"

			err = k8sClient.Update(ctx, apiDefinition)
			Expect(err).ToNot(HaveOccurred())

			By("Check API definition processing status")
			Eventually(func() bool {
				k8sErr := k8sClient.Get(ctx, apiLookupKey, savedApiDefinition)
				return k8sErr == nil && savedApiDefinition.Status.ProcessingStatus == gio.ProcessingStatusFailed
			}, timeout, interval).Should(BeTrue())

			By("Check events")
			Expect(getEventsReason(apiDefinitionFixture)).Should(ContainElements([]string{"Failed"}))
			Expect(getEventsReason(apiDefinitionFixture)).ShouldNot(ContainElements([]string{"Updated"}))

			By("Set right credentials in ManagementContext")
			managementContextRight := managementContextBad.DeepCopy()
			managementContextRight.Spec = managementContextFixture.Spec

			err = k8sClient.Update(ctx, managementContextRight)
			Expect(err).ToNot(HaveOccurred())

			By("Check that API definition has been reconciled on ManagementContext update")

			apimClient, err := internal.NewApimClient(ctx)
			Expect(err).ToNot(HaveOccurred())

			Eventually(func() bool {
				api, apiErr := apimClient.GetByCrossId(apiDefinition.Status.CrossID)
				return apiErr == nil && api.Name == "new-name" && api.Id == apiDefinition.Status.ID
			}, timeout, interval).Should(BeTrue())

			By("Check API definition processing status")
			Eventually(func() bool {
				k8sErr := k8sClient.Get(ctx, apiLookupKey, savedApiDefinition)
				return k8sErr == nil &&
					savedApiDefinition.Status.ProcessingStatus == gio.ProcessingStatusCompleted &&
					k8sUtil.ContainsFinalizer(savedApiDefinition, keys.ApiDefinitionDeletionFinalizer)
			}, timeout, interval).Should(BeTrue())

			By("Check events")
			Expect(getEventsReason(apiDefinitionFixture)).Should(ContainElements([]string{"Updated"}))
		})

		It("Should requeue reconcile with bad ManagementContext BaseUrl", func() {

			By("Set bad BaseUrl in ManagementContext")
			managementContextBad := managementContextFixture.DeepCopy()
			managementContextBad.Spec.BaseUrl = "http://bad-url:8083"

			err := k8sClient.Update(ctx, managementContextBad)
			Expect(err).ToNot(HaveOccurred())

			By("Update the API definition")
			apiDefinition := savedApiDefinition.DeepCopy()
			apiDefinition.Spec.Name = "new-name"

			err = k8sClient.Update(ctx, apiDefinition)
			Expect(err).ToNot(HaveOccurred())

			By("Check API definition processing status")
			Eventually(func() bool {
				k8sErr := k8sClient.Get(ctx, apiLookupKey, savedApiDefinition)
				return k8sErr == nil && savedApiDefinition.Status.ProcessingStatus == gio.ProcessingStatusReconciling
			}, timeout, interval).Should(BeTrue())

			By("Set right BaseUrl in ManagementContext")
			managementContextRight := managementContextBad.DeepCopy()
			managementContextRight.Spec = managementContextFixture.Spec

			err = k8sClient.Update(ctx, managementContextRight)
			Expect(err).ToNot(HaveOccurred())

			By("Check events")
			Expect(getEventsReason(apiDefinitionFixture)).Should(ContainElements([]string{"Reconciling"}))

			By("Check API definition processing status")
			Eventually(func() bool {
				k8sErr := k8sClient.Get(ctx, apiLookupKey, savedApiDefinition)
				return k8sErr == nil && savedApiDefinition.Status.ProcessingStatus == gio.ProcessingStatusCompleted
			}, timeout, interval).Should(BeTrue())
		})
	})
})
