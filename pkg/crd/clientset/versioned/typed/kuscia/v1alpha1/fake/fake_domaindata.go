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

// FakeDomainDatas implements DomainDataInterface
type FakeDomainDatas struct {
	Fake *FakeKusciaV1alpha1
	ns   string
}

var domaindatasResource = schema.GroupVersionResource{Group: "kuscia.secretflow", Version: "v1alpha1", Resource: "domaindatas"}

var domaindatasKind = schema.GroupVersionKind{Group: "kuscia.secretflow", Version: "v1alpha1", Kind: "DomainData"}

// Get takes name of the domainData, and returns the corresponding domainData object, and an error if there is any.
func (c *FakeDomainDatas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DomainData, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domaindatasResource, c.ns, name), &v1alpha1.DomainData{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainData), err
}

// List takes label and field selectors, and returns the list of DomainDatas that match those selectors.
func (c *FakeDomainDatas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DomainDataList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domaindatasResource, domaindatasKind, c.ns, opts), &v1alpha1.DomainDataList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainDataList{ListMeta: obj.(*v1alpha1.DomainDataList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainDataList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domainDatas.
func (c *FakeDomainDatas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domaindatasResource, c.ns, opts))

}

// Create takes the representation of a domainData and creates it.  Returns the server's representation of the domainData, and an error, if there is any.
func (c *FakeDomainDatas) Create(ctx context.Context, domainData *v1alpha1.DomainData, opts v1.CreateOptions) (result *v1alpha1.DomainData, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domaindatasResource, c.ns, domainData), &v1alpha1.DomainData{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainData), err
}

// Update takes the representation of a domainData and updates it. Returns the server's representation of the domainData, and an error, if there is any.
func (c *FakeDomainDatas) Update(ctx context.Context, domainData *v1alpha1.DomainData, opts v1.UpdateOptions) (result *v1alpha1.DomainData, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domaindatasResource, c.ns, domainData), &v1alpha1.DomainData{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainData), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomainDatas) UpdateStatus(ctx context.Context, domainData *v1alpha1.DomainData, opts v1.UpdateOptions) (*v1alpha1.DomainData, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domaindatasResource, "status", c.ns, domainData), &v1alpha1.DomainData{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainData), err
}

// Delete takes name of the domainData and deletes it. Returns an error if one occurs.
func (c *FakeDomainDatas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(domaindatasResource, c.ns, name, opts), &v1alpha1.DomainData{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomainDatas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domaindatasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainDataList{})
	return err
}

// Patch applies the patch and returns the patched domainData.
func (c *FakeDomainDatas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DomainData, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domaindatasResource, c.ns, name, pt, data, subresources...), &v1alpha1.DomainData{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainData), err
}
