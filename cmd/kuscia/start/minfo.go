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

package start

import (
	"context"
	"sync"

	"github.com/secretflow/kuscia/cmd/kuscia/modules"
	"github.com/secretflow/kuscia/pkg/common"
)

type NewModuleFunc func(*modules.ModuleRuntimeConfigs) (modules.Module, error)

type moduleInfo struct {
	name    string
	creator NewModuleFunc
	modes   map[common.RunModeType]bool

	// dependency modules(I depend on others)
	dependencies []string

	// runtime
	instance modules.Module
	ctx      context.Context
	cancel   context.CancelFunc

	isReadyWaitDone bool
	readyError      error

	// run finished?
	finishWG sync.WaitGroup
}

// all dependencies are started and ready now
func (mc *moduleInfo) isModuleDepReady(modules map[string]*moduleInfo) bool {
	startModule := true
	for _, dp := range mc.dependencies {
		if m, ok := modules[dp]; ok {
			if m.instance == nil || !m.isReadyWaitDone {
				startModule = false
				break
			}
		}
	}

	return startModule
}
