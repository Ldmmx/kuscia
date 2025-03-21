/*
Copyright 2016 The Kubernetes Authors.

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

package kri

import (
	"testing"

	"k8s.io/apimachinery/pkg/types"

	pkgcontainer "github.com/secretflow/kuscia/pkg/agent/container"
)

func TestReasonCache(t *testing.T) {
	// Create test sync result
	syncResult := pkgcontainer.PodSyncResult{}
	results := []*pkgcontainer.SyncResult{
		// reason cache should be set for SyncResult with StartContainer action and error
		pkgcontainer.NewSyncResult(pkgcontainer.StartContainer, "container_1"),
		// reason cache should not be set for SyncResult with StartContainer action but without error
		pkgcontainer.NewSyncResult(pkgcontainer.StartContainer, "container_2"),
		// reason cache should not be set for SyncResult with other actions
		pkgcontainer.NewSyncResult(pkgcontainer.KillContainer, "container_3"),
	}
	results[0].Fail(pkgcontainer.ErrRunContainer, "message_1")
	results[2].Fail(pkgcontainer.ErrKillContainer, "message_3")
	syncResult.AddSyncResult(results...)
	uid := types.UID("pod_1")

	reasonCache := NewReasonCache()
	reasonCache.Update(uid, syncResult)
	assertReasonInfo(t, reasonCache, uid, results[0], true)
	assertReasonInfo(t, reasonCache, uid, results[1], false)
	assertReasonInfo(t, reasonCache, uid, results[2], false)

	reasonCache.Remove(uid, results[0].Target.(string))
	assertReasonInfo(t, reasonCache, uid, results[0], false)
}

func assertReasonInfo(t *testing.T, cache *ReasonCache, uid types.UID, result *pkgcontainer.SyncResult, found bool) {
	name := result.Target.(string)
	actualReason, ok := cache.Get(uid, name)
	if ok && !found {
		t.Fatalf("unexpected cache hit: %v, %q", actualReason.Err, actualReason.Message)
	}
	if !ok && found {
		t.Fatalf("corresponding reason info not found")
	}
	if !found {
		return
	}
	reason := result.Error
	message := result.Message
	if actualReason.Err != reason || actualReason.Message != message {
		t.Errorf("expected %v %q, got %v %q", reason, message, actualReason.Err, actualReason.Message)
	}
}
