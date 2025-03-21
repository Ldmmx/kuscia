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
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	kusciaapisv1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	kusciaclientsetfake "github.com/secretflow/kuscia/pkg/crd/clientset/versioned/fake"
	kusciainformers "github.com/secretflow/kuscia/pkg/crd/informers/externalversions"
	"github.com/secretflow/kuscia/pkg/interconn/bfia/common"
	"github.com/secretflow/kuscia/pkg/web/api"
	"github.com/secretflow/kuscia/pkg/web/errorcode"
	"github.com/secretflow/kuscia/proto/api/v1/interconn"
)

func TestNewStartJobHandler(t *testing.T) {
	ctx := context.Background()
	kusciaFakeClient := kusciaclientsetfake.NewSimpleClientset()
	rm, _ := NewResourcesManager(ctx, kusciaFakeClient)

	tests := []struct {
		name       string
		wantNotNil bool
	}{
		{
			name:       "new start job handler",
			wantNotNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantNotNil, NewStartJobHandler(rm) != nil)
		})
	}
}

func Test_getFailedMessageFromKusciaJob(t *testing.T) {
	kj := makeKusciaJob("job-1", nil, kusciaapisv1alpha1.JobStartStage)
	kj.Status.Phase = kusciaapisv1alpha1.KusciaJobFailed
	type args struct {
		kj *kusciaapisv1alpha1.KusciaJob
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get failed message",
			args: args{
				kj: kj,
			},
			want: common.ErrJobStatusFailed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getFailedMessageFromKusciaJob(tt.args.kj))
		})
	}
}

func Test_startJobHandler_GetType(t *testing.T) {
	ctx := context.Background()
	kusciaFakeClient := kusciaclientsetfake.NewSimpleClientset()
	rm, _ := NewResourcesManager(ctx, kusciaFakeClient)
	h := NewStartJobHandler(rm)

	tests := []struct {
		name         string
		wantReqType  reflect.Type
		wantRespType reflect.Type
	}{
		{
			name:         "get req and resp type",
			wantReqType:  reflect.TypeOf(interconn.StartJobRequest{}),
			wantRespType: reflect.TypeOf(interconn.CommonResponse{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReqType, gotRespType := h.GetType()
			assert.Equal(t, tt.wantReqType, gotReqType)
			assert.Equal(t, tt.wantRespType, gotRespType)
		})
	}
}

func Test_startJobHandler_Handle(t *testing.T) {
	kusciaFakeClient := kusciaclientsetfake.NewSimpleClientset()
	kusciaInformerFactory := kusciainformers.NewSharedInformerFactory(kusciaFakeClient, 0)
	kjInformer := kusciaInformerFactory.Kuscia().V1alpha1().KusciaJobs()

	kj1 := makeKusciaJob("job-1", nil, "")
	kj2 := makeKusciaJob("job-2", nil, kusciaapisv1alpha1.JobStartStage)
	kjInformer.Informer().GetStore().Add(kj1)
	kjInformer.Informer().GetStore().Add(kj2)

	rm := &ResourcesManager{
		KusciaClient: kusciaFakeClient,
		KjLister:     kjInformer.Lister(),
		jobTaskInfo:  make(map[string]map[string]struct{}),
		taskJobInfo:  make(map[string]string),
	}
	h := NewStartJobHandler(rm)

	type args struct {
		ctx     *api.BizContext
		request api.ProtoRequest
	}

	tests := []struct {
		name     string
		args     args
		wantCode int32
	}{
		{
			name: "kuscia job doesn't exist",
			args: args{
				ctx: nil,
				request: &interconn.StartJobRequest{
					JobId: "job-11",
				},
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "update kuscia job failed",
			args: args{
				ctx: nil,
				request: &interconn.StartJobRequest{
					JobId: "job-1",
				},
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "update kuscia job succeeded",
			args: args{
				ctx: nil,
				request: &interconn.StartJobRequest{
					JobId: "job-2",
				},
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := h.Handle(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantCode, resp.(*interconn.CommonResponse).Code)
		})
	}
}

func Test_startJobHandler_Validate(t *testing.T) {
	ctx := context.Background()
	kusciaFakeClient := kusciaclientsetfake.NewSimpleClientset()
	rm, _ := NewResourcesManager(ctx, kusciaFakeClient)
	h := NewStartJobHandler(rm)

	type args struct {
		ctx     *api.BizContext
		request api.ProtoRequest
		errs    *errorcode.Errs
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "req type is invalid",
			args: args{
				ctx:     nil,
				request: &interconn.StopJobRequest{},
				errs:    &errorcode.Errs{},
			},
			wantErr: true,
		},
		{
			name: "req job id is empty",
			args: args{
				ctx:     nil,
				request: &interconn.StartJobRequest{},
				errs:    &errorcode.Errs{},
			},
			wantErr: true,
		},
		{
			name: "req is valid",
			args: args{
				ctx: nil,
				request: &interconn.StartJobRequest{
					JobId: "job-1",
				},
				errs: &errorcode.Errs{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.Validate(tt.args.ctx, tt.args.request, tt.args.errs)
			assert.Equal(t, tt.wantErr, len(*tt.args.errs) != 0)
		})
	}
}
