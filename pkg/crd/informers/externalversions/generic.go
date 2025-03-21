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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=kuscia.secretflow, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("appimages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().AppImages().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterdomainroutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().ClusterDomainRoutes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("domains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().Domains().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("domainappimages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().DomainAppImages().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("domaindatas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().DomainDatas().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("domaindatagrants"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().DomainDataGrants().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("domaindatasources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().DomainDataSources().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("domainroutes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().DomainRoutes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().Gateways().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("interopconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().InteropConfigs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("kusciadeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().KusciaDeployments().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("kusciadeploymentsummaries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().KusciaDeploymentSummaries().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("kusciajobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().KusciaJobs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("kusciajobsummaries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().KusciaJobSummaries().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("kusciatasks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().KusciaTasks().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("kusciatasksummaries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().KusciaTaskSummaries().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("taskresources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().TaskResources().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("taskresourcegroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kuscia().V1alpha1().TaskResourceGroups().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
