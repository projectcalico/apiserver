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

// FakeGlobalNetworkPolicies implements GlobalNetworkPolicyInterface
type FakeGlobalNetworkPolicies struct {
	Fake *FakeProjectcalicoV3
}

var globalnetworkpoliciesResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "globalnetworkpolicies"}

var globalnetworkpoliciesKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "GlobalNetworkPolicy"}

// Get takes name of the globalNetworkPolicy, and returns the corresponding globalNetworkPolicy object, and an error if there is any.
func (c *FakeGlobalNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(globalnetworkpoliciesResource, name), &v3.GlobalNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of GlobalNetworkPolicies that match those selectors.
func (c *FakeGlobalNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(globalnetworkpoliciesResource, globalnetworkpoliciesKind, opts), &v3.GlobalNetworkPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.GlobalNetworkPolicyList{ListMeta: obj.(*v3.GlobalNetworkPolicyList).ListMeta}
	for _, item := range obj.(*v3.GlobalNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalNetworkPolicies.
func (c *FakeGlobalNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(globalnetworkpoliciesResource, opts))
}

// Create takes the representation of a globalNetworkPolicy and creates it.  Returns the server's representation of the globalNetworkPolicy, and an error, if there is any.
func (c *FakeGlobalNetworkPolicies) Create(ctx context.Context, globalNetworkPolicy *v3.GlobalNetworkPolicy, opts v1.CreateOptions) (result *v3.GlobalNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(globalnetworkpoliciesResource, globalNetworkPolicy), &v3.GlobalNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalNetworkPolicy), err
}

// Update takes the representation of a globalNetworkPolicy and updates it. Returns the server's representation of the globalNetworkPolicy, and an error, if there is any.
func (c *FakeGlobalNetworkPolicies) Update(ctx context.Context, globalNetworkPolicy *v3.GlobalNetworkPolicy, opts v1.UpdateOptions) (result *v3.GlobalNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(globalnetworkpoliciesResource, globalNetworkPolicy), &v3.GlobalNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalNetworkPolicy), err
}

// Delete takes name of the globalNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeGlobalNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(globalnetworkpoliciesResource, name), &v3.GlobalNetworkPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(globalnetworkpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.GlobalNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched globalNetworkPolicy.
func (c *FakeGlobalNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(globalnetworkpoliciesResource, name, pt, data, subresources...), &v3.GlobalNetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GlobalNetworkPolicy), err
}
