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

// Modified by Ant Group in 2023.

package source

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/types"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"

	"github.com/secretflow/kuscia/pkg/agent/config"
)

func createFileSourceCfg(namespace string, nodeName types.NodeName, path string, period time.Duration) *InitConfig {
	cfg := &InitConfig{
		Namespace: namespace,
		NodeName:  nodeName,
		SourceCfg: &config.SourceCfg{
			File: config.FileSourceCfg{
				Enable: true,
				Path:   path,
				Period: period,
			},
		},
	}

	return cfg
}

func TestExtractFromBadDataFile(t *testing.T) {
	dirName, err := mkTempDir("file-test")
	if err != nil {
		t.Fatalf("unable to create temp dir: %v", err)
	}
	defer removeAll(dirName, t)

	fileName := filepath.Join(dirName, "test_pod_config")
	err = ioutil.WriteFile(fileName, []byte{1, 2, 3}, 0555)
	if err != nil {
		t.Fatalf("unable to write test file %#v", err)
	}

	ch := make(chan kubetypes.PodUpdate, 1)
	lw := newSourceFile(createFileSourceCfg("test-namespace", "localhost", fileName, time.Millisecond), ch)

	err = lw.listConfig()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	expectEmptyChannel(t, ch)
}

func TestExtractFromEmptyDir(t *testing.T) {
	dirName, err := mkTempDir("file-test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer removeAll(dirName, t)

	ch := make(chan kubetypes.PodUpdate, 1)
	lw := newSourceFile(createFileSourceCfg("test-namespace", "localhost", dirName, time.Millisecond), ch)

	err = lw.listConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	update, ok := <-ch
	if !ok {
		t.Fatalf("unexpected type: %#v", update)
	}
	expected := CreatePodUpdate(kubetypes.SET, kubetypes.FileSource)
	if !apiequality.Semantic.DeepEqual(expected, update) {
		t.Fatalf("expected %#v, got %#v", expected, update)
	}
}

func mkTempDir(prefix string) (string, error) {
	return os.MkdirTemp(os.TempDir(), prefix)
}

func removeAll(dir string, t *testing.T) {
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("unable to remove dir %s: %v", dir, err)
	}
}
