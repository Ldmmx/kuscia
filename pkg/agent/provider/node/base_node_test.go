// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package node

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/secretflow/kuscia/pkg/agent/config"
)

func TestBaseNode_configureCommonNode(t *testing.T) {
	agentConfig := config.DefaultStaticAgentConfig()
	agentConfig.RootDir = "."
	capacityManager, err := NewCapacityManager(config.ContainerRuntime, &agentConfig.Capacity, nil, ".", true)
	assert.NoError(t, err)
	dep := &BaseNodeDependence{
		Runtime:         config.ContainerRuntime,
		Namespace:       "test-namespace",
		Address:         "1.1.1.1",
		CapacityManager: capacityManager,
	}
	n := newBaseNode(dep)
	node := n.configureCommonNode(context.Background(), "test-name")
	assert.Equal(t, "test-name", node.Name)
	assert.Equal(t, n.runtime, node.Labels[labelRuntime])
}
