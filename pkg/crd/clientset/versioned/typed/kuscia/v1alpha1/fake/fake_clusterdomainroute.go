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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterDomainRoutes implements ClusterDomainRouteInterface
type FakeClusterDomainRoutes struct {
	Fake *FakeKusciaV1alpha1
}

var clusterdomainroutesResource = schema.GroupVersionResource{Group: "kuscia.secretflow", Version: "v1alpha1", Resource: "clusterdomainroutes"}

var clusterdomainroutesKind = schema.GroupVersionKind{Group: "kuscia.secretflow", Version: "v1alpha1", Kind: "ClusterDomainRoute"}

// Get takes name of the clusterDomainRoute, and returns the corresponding clusterDomainRoute object, and an error if there is any.
func (c *FakeClusterDomainRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterDomainRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterdomainroutesResource, name), &v1alpha1.ClusterDomainRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainRoute), err
}

// List takes label and field selectors, and returns the list of ClusterDomainRoutes that match those selectors.
func (c *FakeClusterDomainRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterDomainRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterdomainroutesResource, clusterdomainroutesKind, opts), &v1alpha1.ClusterDomainRouteList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterDomainRouteList{ListMeta: obj.(*v1alpha1.ClusterDomainRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterDomainRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterDomainRoutes.
func (c *FakeClusterDomainRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterdomainroutesResource, opts))
}

// Create takes the representation of a clusterDomainRoute and creates it.  Returns the server's representation of the clusterDomainRoute, and an error, if there is any.
func (c *FakeClusterDomainRoutes) Create(ctx context.Context, clusterDomainRoute *v1alpha1.ClusterDomainRoute, opts v1.CreateOptions) (result *v1alpha1.ClusterDomainRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterdomainroutesResource, clusterDomainRoute), &v1alpha1.ClusterDomainRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainRoute), err
}

// Update takes the representation of a clusterDomainRoute and updates it. Returns the server's representation of the clusterDomainRoute, and an error, if there is any.
func (c *FakeClusterDomainRoutes) Update(ctx context.Context, clusterDomainRoute *v1alpha1.ClusterDomainRoute, opts v1.UpdateOptions) (result *v1alpha1.ClusterDomainRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterdomainroutesResource, clusterDomainRoute), &v1alpha1.ClusterDomainRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterDomainRoutes) UpdateStatus(ctx context.Context, clusterDomainRoute *v1alpha1.ClusterDomainRoute, opts v1.UpdateOptions) (*v1alpha1.ClusterDomainRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterdomainroutesResource, "status", clusterDomainRoute), &v1alpha1.ClusterDomainRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainRoute), err
}

// Delete takes name of the clusterDomainRoute and deletes it. Returns an error if one occurs.
func (c *FakeClusterDomainRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterdomainroutesResource, name, opts), &v1alpha1.ClusterDomainRoute{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterDomainRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterdomainroutesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterDomainRouteList{})
	return err
}

// Patch applies the patch and returns the patched clusterDomainRoute.
func (c *FakeClusterDomainRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterDomainRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterdomainroutesResource, name, pt, data, subresources...), &v1alpha1.ClusterDomainRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainRoute), err
}
