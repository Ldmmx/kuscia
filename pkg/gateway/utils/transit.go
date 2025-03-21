// Copyright 2024 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
)

const (
	HeaderTransitHash = "Kuscia-Transit-Hash"
	HeaderTransitFlag = "Kuscia-Transit-Flag"
)

func IsTransit(transit *v1alpha1.Transit) bool {
	return transit != nil && (transit.TransitMethod == v1alpha1.TransitMethodReverseTunnel || transit.TransitMethod == v1alpha1.TransitMethodThirdDomain)
}

func IsThirdPartyTransit(transit *v1alpha1.Transit) bool {
	return transit != nil && (transit.TransitMethod == "" || transit.TransitMethod == v1alpha1.TransitMethodThirdDomain)
}

func IsReverseTunnelTransit(transit *v1alpha1.Transit) bool {
	return transit != nil && transit.TransitMethod == v1alpha1.TransitMethodReverseTunnel
}
