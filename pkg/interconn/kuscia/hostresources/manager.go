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

package hostresources

import (
	"fmt"
	"sync"

	kusciaclientset "github.com/secretflow/kuscia/pkg/crd/clientset/versioned"
	kuscialistersv1alpha1 "github.com/secretflow/kuscia/pkg/crd/listers/kuscia/v1alpha1"
	"github.com/secretflow/kuscia/pkg/utils/nlog"
)

// ResourcesManager is an interface to manage host resources.
type ResourcesManager interface {
	Register(host, member string) error
	Deregister(host, member string)
	SetWorkers(worker int)
	GetHostResourceAccessor(host, member string) ResourcesAccessor
	Stop()
	HasSynced() bool
}

// Options defines some options for host resources manager.
type Options struct {
	MemberKusciaClient          kusciaclientset.Interface
	MemberDomainLister          kuscialistersv1alpha1.DomainLister
	MemberDeploymentLister      kuscialistersv1alpha1.KusciaDeploymentLister
	MemberJobLister             kuscialistersv1alpha1.KusciaJobLister
	MemberTaskLister            kuscialistersv1alpha1.KusciaTaskLister
	MemberTaskResourceLister    kuscialistersv1alpha1.TaskResourceLister
	MemberDomainDataLister      kuscialistersv1alpha1.DomainDataLister
	MemberDomainDataGrantLister kuscialistersv1alpha1.DomainDataGrantLister
}

// hostResourcesManager is used to manage host resources controllers.
type hostResourcesManager struct {
	opts    *Options
	workers int

	mu sync.RWMutex
	// key: host/member
	resourceControllers map[string]*hostResourcesController
}

func NewHostResourcesManager(opts *Options) ResourcesManager {
	return &hostResourcesManager{
		opts:                opts,
		resourceControllers: make(map[string]*hostResourcesController),
	}
}

// Stop is used to stop the host resource manager.
func (m *hostResourcesManager) Stop() {
	nlog.Info("Stop host resource manager")
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, hrc := range m.resourceControllers {
		hrc.stop()
	}
	m.resourceControllers = nil
}

// Register is used to register host resource controller.
func (m *hostResourcesManager) Register(host, member string) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	defer func() {
		if err != nil {
			nlog.Errorf("Register host resources for host/member %v/%v failed", host, member)
			m.Deregister(host, member)
		}
	}()

	if _, exist := m.resourceControllers[generateResourceControllerKey(host, member)]; exist {
		nlog.Infof("%v/%v resource controller is already exist, skip register it...", host, member)
		return nil
	}

	opts := &hostResourcesControllerOptions{
		host:                        host,
		member:                      member,
		memberKusciaClient:          m.opts.MemberKusciaClient,
		memberDomainLister:          m.opts.MemberDomainLister,
		memberDeploymentLister:      m.opts.MemberDeploymentLister,
		memberJobLister:             m.opts.MemberJobLister,
		memberTaskLister:            m.opts.MemberTaskLister,
		memberTaskResourceLister:    m.opts.MemberTaskResourceLister,
		memberDomainDataLister:      m.opts.MemberDomainDataLister,
		memberDomainDataGrantLister: m.opts.MemberDomainDataGrantLister,
	}

	var hrc *hostResourcesController
	hrc, err = newHostResourcesController(opts)
	if err != nil {
		return fmt.Errorf("register host resource controller failed, %v", err.Error())
	}

	m.resourceControllers[generateResourceControllerKey(host, member)] = hrc
	go func() {
		err = hrc.run(m.workers)
		if err != nil {
			nlog.Errorf("run resource controller failed, %v", err.Error())
		}
	}()
	return
}

// Deregister is used to deregister host resource controller.
func (m *hostResourcesManager) Deregister(host, member string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	hrc, exist := m.resourceControllers[generateResourceControllerKey(host, member)]
	if !exist {
		return
	}

	hrc.stop()
	delete(m.resourceControllers, generateResourceControllerKey(host, member))
	nlog.Infof("Deregister host/member %v/%v resources controller", host, member)
}

// SetWorkers is used to set workers for host resource manager.
func (m *hostResourcesManager) SetWorkers(workers int) {
	m.workers = workers
}

// GetHostResourceAccessor is used to get host resource accessor.
func (m *hostResourcesManager) GetHostResourceAccessor(host, member string) ResourcesAccessor {
	m.mu.RLock()
	defer m.mu.RUnlock()
	ra, exist := m.resourceControllers[generateResourceControllerKey(host, member)]
	if !exist {
		return nil
	}
	return ra
}

// HasSynced is used to check if there are already resources that have been synchronized.
func (m *hostResourcesManager) HasSynced() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if len(m.resourceControllers) == 0 {
		return true
	}

	for _, rc := range m.resourceControllers {
		if rc.HasSynced() {
			return true
		}
	}
	return false
}

// generateResourceControllerKey is used to generate resource controller key.
func generateResourceControllerKey(host, member string) string {
	return host + "/" + member
}
