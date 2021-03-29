// Copyright (c) 2021 Tigera, Inc. All rights reserved.
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

	v3 "github.com/projectcalico/apiserver/pkg/apis/projectcalico/v3"
)

// FakeNetworkSets implements NetworkSetInterface
type FakeNetworkSets struct {
	Fake *FakeProjectcalicoV3
	ns   string
}

var networksetsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "networksets"}

var networksetsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "NetworkSet"}

// Get takes name of the networkSet, and returns the corresponding networkSet object, and an error if there is any.
func (c *FakeNetworkSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networksetsResource, c.ns, name), &v3.NetworkSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NetworkSet), err
}

// List takes label and field selectors, and returns the list of NetworkSets that match those selectors.
func (c *FakeNetworkSets) List(ctx context.Context, opts v1.ListOptions) (result *v3.NetworkSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networksetsResource, networksetsKind, c.ns, opts), &v3.NetworkSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.NetworkSetList{ListMeta: obj.(*v3.NetworkSetList).ListMeta}
	for _, item := range obj.(*v3.NetworkSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkSets.
func (c *FakeNetworkSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networksetsResource, c.ns, opts))

}

// Create takes the representation of a networkSet and creates it.  Returns the server's representation of the networkSet, and an error, if there is any.
func (c *FakeNetworkSets) Create(ctx context.Context, networkSet *v3.NetworkSet, opts v1.CreateOptions) (result *v3.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networksetsResource, c.ns, networkSet), &v3.NetworkSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NetworkSet), err
}

// Update takes the representation of a networkSet and updates it. Returns the server's representation of the networkSet, and an error, if there is any.
func (c *FakeNetworkSets) Update(ctx context.Context, networkSet *v3.NetworkSet, opts v1.UpdateOptions) (result *v3.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networksetsResource, c.ns, networkSet), &v3.NetworkSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NetworkSet), err
}

// Delete takes name of the networkSet and deletes it. Returns an error if one occurs.
func (c *FakeNetworkSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networksetsResource, c.ns, name), &v3.NetworkSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networksetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.NetworkSetList{})
	return err
}

// Patch applies the patch and returns the patched networkSet.
func (c *FakeNetworkSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.NetworkSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networksetsResource, c.ns, name, pt, data, subresources...), &v3.NetworkSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NetworkSet), err
}
