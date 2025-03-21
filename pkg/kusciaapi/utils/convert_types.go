// Copyright 2023 Ant Group Co., Ltd.
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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Int32Value(v *int) int32 {
	if v != nil {
		return int32(*v)
	}
	return 0
}

func IntValue(v int32) *int {
	i := int(v)
	return &i
}

func TimeRfc3339String(t *metav1.Time) string {
	if t != nil {
		if t.IsZero() {
			// Encode unset/nil objects as an empty string
			return ""
		}
		return t.UTC().Format(time.RFC3339)
	}
	return ""
}
