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

package layout

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBundle(t *testing.T) {
	rootDir := t.TempDir()
	bundle, err := NewBundle(rootDir)
	assert.NoError(t, err)

	containerBundle := bundle.GetContainerBundle("aaa")
	assert.Equal(t, filepath.Join(rootDir, "aaa"), containerBundle.GetRootDirectory())
	assert.Equal(t, bundleOciRootfsDir, containerBundle.GetOciRootfsName())
	assert.Equal(t, filepath.Join(rootDir, "aaa", bundleWorkingDir), containerBundle.GetFsWorkingDirPath())
	assert.Equal(t, filepath.Join(rootDir, "aaa", bundleOciConfigFile), containerBundle.GetOciConfigPath())
	assert.Equal(t, filepath.Join(rootDir, "aaa", bundleOciRootfsDir), containerBundle.GetOciRootfsPath())
}
