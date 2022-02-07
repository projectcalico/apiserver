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

// FakeGlobalNetworkSets implements GlobalNetworkSetInterface
type FakeGlobalNetworkSets struct {
	Fake *FakeProjectcalico
}

var globalnetworksetsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "globalnetworksets"}

var globalnetworksetsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "GlobalNetworkSet"}

// Get takes name of the globalNetworkSet, and returns the corresponding globalNetworkSet object, and an error if there is any.
func (c *FakeGlobalNetworkSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.GlobalNetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(globalnetworksetsResource, name), &projectcalico.GlobalNetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalNetworkSet), err
}

// List takes label and field selectors, and returns the list of GlobalNetworkSets that match those selectors.
func (c *FakeGlobalNetworkSets) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.GlobalNetworkSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(globalnetworksetsResource, globalnetworksetsKind, opts), &projectcalico.GlobalNetworkSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.GlobalNetworkSetList{ListMeta: obj.(*projectcalico.GlobalNetworkSetList).ListMeta}
	for _, item := range obj.(*projectcalico.GlobalNetworkSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalNetworkSets.
func (c *FakeGlobalNetworkSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(globalnetworksetsResource, opts))
}

// Create takes the representation of a globalNetworkSet and creates it.  Returns the server's representation of the globalNetworkSet, and an error, if there is any.
func (c *FakeGlobalNetworkSets) Create(ctx context.Context, globalNetworkSet *projectcalico.GlobalNetworkSet, opts v1.CreateOptions) (result *projectcalico.GlobalNetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(globalnetworksetsResource, globalNetworkSet), &projectcalico.GlobalNetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalNetworkSet), err
}

// Update takes the representation of a globalNetworkSet and updates it. Returns the server's representation of the globalNetworkSet, and an error, if there is any.
func (c *FakeGlobalNetworkSets) Update(ctx context.Context, globalNetworkSet *projectcalico.GlobalNetworkSet, opts v1.UpdateOptions) (result *projectcalico.GlobalNetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(globalnetworksetsResource, globalNetworkSet), &projectcalico.GlobalNetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalNetworkSet), err
}

// Delete takes name of the globalNetworkSet and deletes it. Returns an error if one occurs.
func (c *FakeGlobalNetworkSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(globalnetworksetsResource, name), &projectcalico.GlobalNetworkSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalNetworkSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(globalnetworksetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &projectcalico.GlobalNetworkSetList{})
	return err
}

// Patch applies the patch and returns the patched globalNetworkSet.
func (c *FakeGlobalNetworkSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.GlobalNetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(globalnetworksetsResource, name, pt, data, subresources...), &projectcalico.GlobalNetworkSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalNetworkSet), err
}
