// Copyright (c) 2021,2022 Tigera, Inc. All rights reserved.
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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
)

// FakeFelixConfigurations implements FelixConfigurationInterface
type FakeFelixConfigurations struct {
	Fake *FakeProjectcalico
}

var felixconfigurationsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "felixconfigurations"}

var felixconfigurationsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "FelixConfiguration"}

// Get takes name of the felixConfiguration, and returns the corresponding felixConfiguration object, and an error if there is any.
func (c *FakeFelixConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.FelixConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(felixconfigurationsResource, name), &projectcalico.FelixConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.FelixConfiguration), err
}

// List takes label and field selectors, and returns the list of FelixConfigurations that match those selectors.
func (c *FakeFelixConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.FelixConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(felixconfigurationsResource, felixconfigurationsKind, opts), &projectcalico.FelixConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.FelixConfigurationList{ListMeta: obj.(*projectcalico.FelixConfigurationList).ListMeta}
	for _, item := range obj.(*projectcalico.FelixConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested felixConfigurations.
func (c *FakeFelixConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(felixconfigurationsResource, opts))
}

// Create takes the representation of a felixConfiguration and creates it.  Returns the server's representation of the felixConfiguration, and an error, if there is any.
func (c *FakeFelixConfigurations) Create(ctx context.Context, felixConfiguration *projectcalico.FelixConfiguration, opts v1.CreateOptions) (result *projectcalico.FelixConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(felixconfigurationsResource, felixConfiguration), &projectcalico.FelixConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.FelixConfiguration), err
}

// Update takes the representation of a felixConfiguration and updates it. Returns the server's representation of the felixConfiguration, and an error, if there is any.
func (c *FakeFelixConfigurations) Update(ctx context.Context, felixConfiguration *projectcalico.FelixConfiguration, opts v1.UpdateOptions) (result *projectcalico.FelixConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(felixconfigurationsResource, felixConfiguration), &projectcalico.FelixConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.FelixConfiguration), err
}

// Delete takes name of the felixConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeFelixConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(felixconfigurationsResource, name), &projectcalico.FelixConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFelixConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(felixconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &projectcalico.FelixConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched felixConfiguration.
func (c *FakeFelixConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.FelixConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(felixconfigurationsResource, name, pt, data, subresources...), &projectcalico.FelixConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.FelixConfiguration), err
}
