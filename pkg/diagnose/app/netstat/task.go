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

package netstat

import (
	"context"

	"github.com/secretflow/kuscia/pkg/diagnose/app/client"
	dcommon "github.com/secretflow/kuscia/pkg/diagnose/common"
	"github.com/secretflow/kuscia/proto/api/v1alpha1/diagnose"
)

type Task interface {
	Run(ctx context.Context)
	Output() *TaskOutput
	Name() string
}

type TaskOutput diagnose.MetricItem

func (t *TaskOutput) ToStringArray() []string {
	return []string{t.Name, t.DetectedValue, t.Threshold, t.Result, t.Information}
}

func OutputsToMetricItems(elements []*TaskOutput) []*diagnose.MetricItem {
	items := make([]*diagnose.MetricItem, 0)
	for _, element := range elements {
		items = append(items, (*diagnose.MetricItem)(element))
	}
	return items
}

func MetricItemsToOutputs(elements []*diagnose.MetricItem) []*TaskOutput {
	items := make([]*TaskOutput, 0)
	for _, element := range elements {
		items = append(items, (*TaskOutput)(element))
	}
	return items
}

type TaskGroup struct {
	tasks          []Task
	diagnoseClient *client.Client
	JobConfig      *NetworkJobConfig
}

type NetworkJobConfig struct {
	ServerPort   int
	SelfDomain   string
	SelfEndpoint string
	PeerDomain   string `json:"peer_domain"`
	PeerEndpoint string `json:"peer_endpoint"`
	Manual       bool   `json:"manual"`
	JobID        string `json:"job_id"`
	*NetworkParam
}

type NetworkParam struct {
	Pass              bool `json:"pass"`
	Speed             bool `json:"speed"`
	SpeedThres        int  `json:"speed_thres"`
	RTT               bool `json:"rtt"`
	RTTTres           int  `json:"rtt_thres"`
	ProxyTimeout      bool `json:"proxy_timeout"`
	ProxyTimeoutThres int  `json:"proxy_timeout_thres"`
	Size              bool `json:"size"`
	SizeThres         int  `json:"size_thres"`
	ProxyBuffer       bool `json:"proxy_buffer"`
	Bidirection       bool `json:"bi_direction"`
}

func NewTaskGroup(cli *client.Client, config *NetworkParam) *TaskGroup {
	tg := new(TaskGroup)
	tg.diagnoseClient = cli

	tg.tasks = append(tg.tasks, NewConnectionTask(tg.diagnoseClient))

	if config.Speed {
		tg.tasks = append(tg.tasks, NewBandWidthTask(tg.diagnoseClient, config.SpeedThres))
	}
	if config.RTT {
		tg.tasks = append(tg.tasks, NewLatencyTask(tg.diagnoseClient, config.RTTTres))
	}
	if config.ProxyTimeout {
		tg.tasks = append(tg.tasks, NewProxyTimeoutTask(tg.diagnoseClient, config.ProxyTimeoutThres))
	}
	if config.Size {
		tg.tasks = append(tg.tasks, NewReqSizeTask(tg.diagnoseClient, config.SizeThres))
	}
	if config.ProxyBuffer {
		tg.tasks = append(tg.tasks, NewBufferTask(tg.diagnoseClient))
	}
	return tg
}

func (tg *TaskGroup) Start(ctx context.Context) ([]*TaskOutput, error) {
	// validate connection task first
	results := make([]*TaskOutput, 0)
	for i, task := range tg.tasks {
		task.Run(ctx)
		results = append(results, task.Output())
		if i == 0 && task.Output().Result == dcommon.Fail {
			break
		}
	}
	return results, nil
}
