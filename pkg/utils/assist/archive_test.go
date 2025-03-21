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

package assist

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTar(t *testing.T) {
	rootDir := t.TempDir()
	srcDir := filepath.Join(rootDir, "src")
	assert.NoError(t, os.MkdirAll(srcDir, 0755))
	assert.NoError(t, os.WriteFile(filepath.Join(srcDir, "a.txt"), []byte("aaa"), 0644))
	assert.NoError(t, os.WriteFile(filepath.Join(srcDir, "b.txt"), []byte("bbb"), 0644))

	dstPath := filepath.Join(rootDir, "dst.tar")
	dst, err := os.Create(dstPath)
	assert.NoError(t, err)

	assert.NoError(t, Tar(srcDir, true, dst))
	targetDir := filepath.Join(rootDir, "target")
	assert.NoError(t, ExtractTarFile(targetDir, dstPath, false, false))
}
