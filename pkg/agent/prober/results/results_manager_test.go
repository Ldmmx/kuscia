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

package results

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	pkgcontainer "github.com/secretflow/kuscia/pkg/agent/container"
)

func TestCacheOperations(t *testing.T) {
	m := NewManager()

	unsetID := pkgcontainer.CtrID{Type: "test", ID: "unset"}
	setID := pkgcontainer.CtrID{Type: "test", ID: "set"}

	_, found := m.Get(unsetID)
	assert.False(t, found, "unset result found")

	m.Set(setID, Success, &v1.Pod{})
	result, found := m.Get(setID)
	assert.True(t, result == Success, "set result")
	assert.True(t, found, "set result found")

	m.Remove(setID)
	_, found = m.Get(setID)
	assert.False(t, found, "removed result found")
}

func TestUpdates(t *testing.T) {
	m := NewManager()

	pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod"}}
	fooID := pkgcontainer.CtrID{Type: "test", ID: "foo"}
	barID := pkgcontainer.CtrID{Type: "test", ID: "bar"}

	expectUpdate := func(expected Update, msg string) {
		select {
		case u := <-m.Updates():
			if expected != u {
				t.Errorf("Expected update %v, received %v: %s", expected, u, msg)
			}
		case <-time.After(wait.ForeverTestTimeout):
			t.Errorf("Timed out waiting for update %v: %s", expected, msg)
		}
	}

	expectNoUpdate := func(msg string) {
		// NOTE: Since updates are accumulated asynchronously, this method is not guaranteed to fail
		// when it should. In the event it misses a failure, the following calls to expectUpdate should
		// still fail.
		select {
		case u := <-m.Updates():
			t.Errorf("Unexpected update %v: %s", u, msg)
		default:
			// Pass
		}
	}

	// New result should always push an update.
	m.Set(fooID, Success, pod)
	expectUpdate(Update{fooID, Success, pod.UID}, "new success")

	m.Set(barID, Failure, pod)
	expectUpdate(Update{barID, Failure, pod.UID}, "new failure")

	// Unchanged results should not send an update.
	m.Set(fooID, Success, pod)
	expectNoUpdate("unchanged foo")

	m.Set(barID, Failure, pod)
	expectNoUpdate("unchanged bar")

	// Changed results should send an update.
	m.Set(fooID, Failure, pod)
	expectUpdate(Update{fooID, Failure, pod.UID}, "changed foo")

	m.Set(barID, Success, pod)
	expectUpdate(Update{barID, Success, pod.UID}, "changed bar")
}
