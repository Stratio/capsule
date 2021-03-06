//+build e2e

/*
Copyright 2020 Clastix Labs.

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

package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	"github.com/clastix/capsule/api/v1alpha1"
)

var _ = Describe("creating a Namespace in over-quota of three", func() {
	tnt := &v1alpha1.Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name: "over-quota-tenant",
		},
		Spec: v1alpha1.TenantSpec{
			Owner: v1alpha1.OwnerSpec{
				Name: "bob",
				Kind: "User",
			},
			NamespaceQuota: pointer.Int32Ptr(3),
		},
	}

	JustBeforeEach(func() {
		EventuallyCreation(func() error {
			return k8sClient.Create(context.TODO(), tnt)
		}).Should(Succeed())
	})
	JustAfterEach(func() {
		Expect(k8sClient.Delete(context.TODO(), tnt)).Should(Succeed())
	})

	It("should fail", func() {
		By("creating three Namespaces", func() {
			for _, name := range []string{"bob-dev", "bob-staging", "bob-production"} {
				ns := NewNamespace(name)
				NamespaceCreation(ns, tnt, defaultTimeoutInterval).Should(Succeed())
				TenantNamespaceList(tnt, podRecreationTimeoutInterval).Should(ContainElement(ns.GetName()))
			}
		})

		ns := NewNamespace("bob-fail")
		cs := ownerClient(tnt)
		_, err := cs.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
		Expect(err).ShouldNot(Succeed())
	})
})
