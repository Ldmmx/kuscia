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

package modules

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RunReporter(t *testing.T) {
	runCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config := mockModuleRuntimeConfig(t)
	m, err := NewReporter(config)
	assert.NoError(t, err)
	assert.NoError(t, runModule(runCtx, m))
}
