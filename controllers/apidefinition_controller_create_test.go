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
package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	model "github.com/gravitee-io/gravitee-kubernetes-operator/api/model"
	gio "github.com/gravitee-io/gravitee-kubernetes-operator/api/v1alpha1"
	managementapi "github.com/gravitee-io/gravitee-kubernetes-operator/internal/apim/managementapi"
	"github.com/gravitee-io/gravitee-kubernetes-operator/internal/utils"
	"github.com/gravitee-io/gravitee-kubernetes-operator/test"
)

var _ = Describe("API Definition Controller", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		namespace = "default"
		origin    = "kubernetes"
		mode      = "fully_managed"

		timeout  = time.Second * 10
		interval = time.Millisecond * 250
	)

	ctx := context.Background()
	httpClient := http.Client{Timeout: 5 * time.Second}

	Context("With basic ApiDefinition", func() {
		var apiDefinitionFixture *gio.ApiDefinition

		var apiLookupKey types.NamespacedName

		BeforeEach(func() {
			By("Without a management context")

			apiDefinition, err := test.NewApiDefinition("../config/samples/apim/basic-example.yml")
			Expect(err).ToNot(HaveOccurred())

			apiDefinitionFixture = apiDefinition
			apiLookupKey = types.NamespacedName{Name: apiDefinitionFixture.Name, Namespace: namespace}
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, apiDefinitionFixture)).Should(Succeed())

			Eventually(func() error {
				return k8sClient.Get(ctx, apiLookupKey, apiDefinitionFixture)
			}, timeout, interval).ShouldNot(Succeed())
		})

		It("Should create an API Definition", func() {
			By("Create an API definition resource without a management context")

			Expect(k8sClient.Create(ctx, apiDefinitionFixture)).Should(Succeed())

			By("Get created resource and expect to find it")

			apiDefinition := new(gio.ApiDefinition)
			Eventually(func() bool {
				err := k8sClient.Get(ctx, apiLookupKey, apiDefinition)
				return err == nil && apiDefinition.Status.CrossID != ""
			}, timeout, interval).Should(BeTrue())

			var endpoint = test.GatewayUrl + apiDefinition.Spec.Proxy.VirtualHosts[0].Path

			expectedApiName := apiDefinitionFixture.Spec.Name
			Expect(apiDefinition.Spec.Name).Should(Equal(expectedApiName))

			By("Call gateway endpoint and expect the API to be available")

			Eventually(func() bool {
				res, callErr := httpClient.Get(endpoint)
				return callErr == nil && res.StatusCode == 200
			}, timeout, interval).Should(BeTrue())
		})
	})

	Context("With basic ApiDefinition & ManagementContext", func() {
		var apiDefinitionFixture *gio.ApiDefinition
		var managementContextFixture *gio.ManagementContext
		var apiLookupKey types.NamespacedName
		var contextLookupKey types.NamespacedName

		BeforeEach(func() {
			managementContext, err := test.NewManagementContext(
				"../config/samples/context/dev/managementcontext_credentials.yaml")
			Expect(err).ToNot(HaveOccurred())

			apiDefinition, err := test.NewApiDefinition("../config/samples/apim/basic-example-with-ctx.yml")
			Expect(err).ToNot(HaveOccurred())

			apiDefinitionFixture = apiDefinition
			managementContextFixture = managementContext
			apiLookupKey = types.NamespacedName{Name: apiDefinitionFixture.Name, Namespace: namespace}
			contextLookupKey = types.NamespacedName{Name: managementContextFixture.Name, Namespace: namespace}
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, apiDefinitionFixture)).Should(Succeed())

			Expect(k8sClient.Delete(ctx, managementContextFixture)).Should(Succeed())

			Eventually(func() error {
				return k8sClient.Get(ctx, apiLookupKey, apiDefinitionFixture)
			}, timeout, interval).ShouldNot(Succeed())

			Eventually(func() error {
				return k8sClient.Get(ctx, contextLookupKey, managementContextFixture)
			}, timeout, interval).ShouldNot(Succeed())
		})

		It("Should create an API Definition", func() {
			By("Create a management context to synchronize with the REST API")
			Expect(k8sClient.Create(ctx, managementContextFixture)).Should(Succeed())

			By("Create an API definition resource referencing the management context")
			Expect(k8sClient.Create(ctx, apiDefinitionFixture)).Should(Succeed())

			By("Get created resource and expect to find it")

			managementContext := new(gio.ManagementContext)
			Eventually(func() error {
				return k8sClient.Get(ctx, contextLookupKey, managementContext)
			}, timeout, interval).Should(Succeed())

			apiDefinition := new(gio.ApiDefinition)
			Eventually(func() bool {
				err := k8sClient.Get(ctx, apiLookupKey, apiDefinition)
				return err == nil && apiDefinition.Status.CrossID != ""
			}, timeout, interval).Should(BeTrue())

			expectedApiName := apiDefinitionFixture.Spec.Name
			Expect(apiDefinition.Spec.Name).Should(Equal(expectedApiName))

			By("Call gateway endpoint and expect the API to be available")

			var endpoint = test.GatewayUrl + apiDefinition.Spec.Proxy.VirtualHosts[0].Path

			Eventually(func() bool {
				res, callErr := httpClient.Get(endpoint)
				return callErr == nil && res.StatusCode == 200
			}, timeout, interval).Should(BeTrue())

			By("Call rest API and expect one API matching status cross ID")

			apimClient := managementapi.NewClient(ctx, managementContextFixture, httpClient)
			Eventually(func() bool {
				api, apiErr := apimClient.GetByCrossId(apiDefinition.Status.CrossID)
				return apiErr == nil && api.Id == apiDefinition.Status.ID
			}, timeout, interval).Should(BeTrue())
		})

		It("Should create a STOPPED API Definition", func() {
			apiDefinitionFixture.Spec.State = model.StateStopped

			By("Create a management context to synchronize with the REST API")
			Expect(k8sClient.Create(ctx, managementContextFixture)).Should(Succeed())

			By("Create an API definition resource referencing the management context")
			Expect(k8sClient.Create(ctx, apiDefinitionFixture)).Should(Succeed())

			By("Get created resource and expect to find it")

			managementContext := new(gio.ManagementContext)
			Eventually(func() error {
				return k8sClient.Get(ctx, contextLookupKey, managementContext)
			}, timeout, interval).Should(Succeed())

			apiDefinition := new(gio.ApiDefinition)
			Eventually(func() bool {
				err := k8sClient.Get(ctx, apiLookupKey, apiDefinition)
				return err == nil && apiDefinition.Status.CrossID != ""
			}, timeout, interval).Should(BeTrue())

			expectedApiName := apiDefinitionFixture.Spec.Name
			Expect(apiDefinition.Spec.Name).Should(Equal(expectedApiName))

			By("Call gateway endpoint and expect the API not to be available")

			var endpoint = test.GatewayUrl + apiDefinition.Spec.Proxy.VirtualHosts[0].Path

			Eventually(func() bool {
				res, callErr := httpClient.Get(endpoint)
				return callErr == nil && res.StatusCode == 404
			}, timeout, interval).Should(BeTrue())

			By("Call rest API and expect one API matching status cross ID and state STOPPED")

			apimClient := managementapi.NewClient(ctx, managementContextFixture, httpClient)
			Eventually(func() bool {
				api, apiErr := apimClient.GetByCrossId(apiDefinition.Status.CrossID)
				return apiErr == nil && api.Id == apiDefinition.Status.ID && api.State == "STOPPED"
			}, timeout, interval).Should(BeTrue())
		})

		It("Should create an API Definition with existing api in Management Api", func() {
			apimClient := managementapi.NewClient(ctx, managementContextFixture, httpClient)

			By("Init existing api in management api")
			existingApiSpec := apiDefinitionFixture.Spec.DeepCopy()
			existingApiSpec.Id = utils.NewUUID()
			existingApiSpec.CrossId = utils.ToUUID(
				types.NamespacedName{Namespace: apiDefinitionFixture.Namespace, Name: apiDefinitionFixture.Name}.String())
			existingApiSpec.DefinitionContext = &model.DefinitionContext{
				Origin: origin,
				Mode:   mode,
			}
			existingApiSpec.Plans = []*model.Plan{
				{
					Id:       utils.ToUUID(existingApiSpec.Id + "/" + "G.K.O. Default"),
					Name:     "G.K.O. Default",
					Security: "KEY_LESS",
					Status:   "PUBLISHED",
				},
			}
			apiJson, err := json.Marshal(existingApiSpec)
			Expect(err).ToNot(HaveOccurred())

			err = apimClient.CreateApi(apiJson)
			Expect(err).ToNot(HaveOccurred())

			By("Create a management context to synchronize with the REST API")
			Expect(k8sClient.Create(ctx, managementContextFixture)).Should(Succeed())

			By("Create an API definition resource referencing the management context")
			Expect(k8sClient.Create(ctx, apiDefinitionFixture)).Should(Succeed())

			By("Get created resource and expect to find it")

			managementContext := new(gio.ManagementContext)
			Eventually(func() error {
				return k8sClient.Get(ctx, contextLookupKey, managementContext)
			}, timeout, interval).Should(Succeed())

			apiDefinition := new(gio.ApiDefinition)
			Eventually(func() bool {
				k8sErr := k8sClient.Get(ctx, apiLookupKey, apiDefinition)
				return k8sErr == nil && apiDefinition.Status.CrossID != ""
			}, timeout, interval).Should(BeTrue())

			expectedApiName := apiDefinitionFixture.Spec.Name
			Expect(apiDefinition.Spec.Name).Should(Equal(expectedApiName))

			By("Call gateway endpoint and expect the API to be available")

			var endpoint = test.GatewayUrl + apiDefinition.Spec.Proxy.VirtualHosts[0].Path

			Eventually(func() bool {
				res, callErr := httpClient.Get(endpoint)
				return callErr == nil && res.StatusCode == 200
			}, timeout, interval).Should(BeTrue())

			By("Call rest API and expect one API matching status cross ID")

			Eventually(func() bool {
				api, apiErr := apimClient.GetByCrossId(apiDefinition.Status.CrossID)
				return apiErr == nil && api.Id == apiDefinition.Status.ID
			}, timeout, interval).Should(BeTrue())
		})

	})
})
