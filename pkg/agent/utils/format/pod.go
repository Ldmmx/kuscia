/*
Copyright 2015 The Kubernetes Authors.

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

// Modified by Ant Group in 2023.

package format

import (
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

// Pod returns a string representing a pod in a consistent human readable format,
// with pod UID as part of the string.
func Pod(pod *v1.Pod) string {
	if pod == nil {
		return "<nil>"
	}
	return PodDesc(pod.Name, pod.Namespace, pod.UID)
}

// Pods returns a string representing pods in a consistent human readable format,
// with pod UID as part of the string.
func Pods(pods []*v1.Pod) string {
	podDescList := make([]string, 0, len(pods))
	for _, pod := range pods {
		podDescList = append(podDescList, Pod(pod))
	}

	return fmt.Sprintf("%v", podDescList)
}

// PodDesc returns a string representing a pod in a consistent human readable format,
// with pod UID as part of the string.
func PodDesc(podName, podNamespace string, podUID types.UID) string {
	// Use underscore as the delimiter because it is not allowed in pod name
	// (DNS subdomain format), while allowed in the container name format.
	return fmt.Sprintf("%s_%s(%s)", podName, podNamespace, podUID)
}

// PodWithDeletionTimestamp is the same as Pod. In addition, it prints the
// deletion timestamp of the pod if it's not nil.
func PodWithDeletionTimestamp(pod *v1.Pod) string {
	if pod == nil {
		return "<nil>"
	}
	var deletionTimestamp string
	if pod.DeletionTimestamp != nil {
		deletionTimestamp = ":DeletionTimestamp=" + pod.DeletionTimestamp.UTC().Format(time.RFC3339)
	}
	return Pod(pod) + deletionTimestamp
}
