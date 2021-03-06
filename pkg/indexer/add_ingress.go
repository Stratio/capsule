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

package indexer

import (
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	networkingv1 "k8s.io/api/networking/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"

	"github.com/clastix/capsule/pkg/indexer/ingress"
	"github.com/clastix/capsule/pkg/webhook/utils"
)

func init() {
	AddToIndexerFuncs = append(AddToIndexerFuncs, ingress.Hostname{Obj: &extensionsv1beta1.Ingress{}})
	// ingresses.networking.k8s.io/v1 introduced by 1.19
	{
		majorVer, minorVer, _, _ := utils.GetK8sVersion()
		if majorVer == 1 && minorVer >= 19 {
			AddToIndexerFuncs = append(AddToIndexerFuncs, ingress.Hostname{Obj: &networkingv1.Ingress{}})
		}
	}
	AddToIndexerFuncs = append(AddToIndexerFuncs, ingress.Hostname{Obj: &networkingv1beta1.Ingress{}})
}
