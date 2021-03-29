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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	v3 "github.com/projectcalico/apiserver/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
)

type FakeProjectcalicoV3 struct {
	*testing.Fake
}

func (c *FakeProjectcalicoV3) BGPConfigurations() v3.BGPConfigurationInterface {
	return &FakeBGPConfigurations{c}
}

func (c *FakeProjectcalicoV3) BGPPeers() v3.BGPPeerInterface {
	return &FakeBGPPeers{c}
}

func (c *FakeProjectcalicoV3) ClusterInformations() v3.ClusterInformationInterface {
	return &FakeClusterInformations{c}
}

func (c *FakeProjectcalicoV3) FelixConfigurations() v3.FelixConfigurationInterface {
	return &FakeFelixConfigurations{c}
}

func (c *FakeProjectcalicoV3) GlobalNetworkPolicies() v3.GlobalNetworkPolicyInterface {
	return &FakeGlobalNetworkPolicies{c}
}

func (c *FakeProjectcalicoV3) GlobalNetworkSets() v3.GlobalNetworkSetInterface {
	return &FakeGlobalNetworkSets{c}
}

func (c *FakeProjectcalicoV3) HostEndpoints() v3.HostEndpointInterface {
	return &FakeHostEndpoints{c}
}

func (c *FakeProjectcalicoV3) IPPools() v3.IPPoolInterface {
	return &FakeIPPools{c}
}

func (c *FakeProjectcalicoV3) KubeControllersConfigurations() v3.KubeControllersConfigurationInterface {
	return &FakeKubeControllersConfigurations{c}
}

func (c *FakeProjectcalicoV3) NetworkPolicies(namespace string) v3.NetworkPolicyInterface {
	return &FakeNetworkPolicies{c, namespace}
}

func (c *FakeProjectcalicoV3) NetworkSets(namespace string) v3.NetworkSetInterface {
	return &FakeNetworkSets{c, namespace}
}

func (c *FakeProjectcalicoV3) Profiles() v3.ProfileInterface {
	return &FakeProfiles{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProjectcalicoV3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
