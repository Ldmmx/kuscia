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

package handler

import (
	"github.com/secretflow/kuscia/pkg/controllers/kusciatask/metrics"
	kusciaapisv1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
)

// SucceededHandler is used to handle kuscia task which phase is Succeeded.
type SucceededHandler struct {
	*FinishedHandler
}

// NewSucceededHandler returns a SucceededHandler instance.
func NewSucceededHandler(finishedHandler *FinishedHandler) *SucceededHandler {
	return &SucceededHandler{
		FinishedHandler: finishedHandler,
	}
}

// Handle is used to perform the real logic.
func (s *SucceededHandler) Handle(kusciaTask *kusciaapisv1alpha1.KusciaTask) (bool, error) {
	needUpdate, err := s.FinishedHandler.Handle(kusciaTask)
	if err != nil {
		return false, nil
	}
	metrics.TaskResultStats.WithLabelValues(metrics.Succeeded).Inc()
	return needUpdate, nil
}
