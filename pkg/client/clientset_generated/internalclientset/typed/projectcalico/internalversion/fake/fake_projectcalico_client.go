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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	internalversion "github.com/projectcalico/apiserver/pkg/client/clientset_generated/internalclientset/typed/projectcalico/internalversion"
)

type FakeProjectcalico struct {
	*testing.Fake
}

func (c *FakeProjectcalico) BGPConfigurations() internalversion.BGPConfigurationInterface {
	return &FakeBGPConfigurations{c}
}

func (c *FakeProjectcalico) BGPPeers() internalversion.BGPPeerInterface {
	return &FakeBGPPeers{c}
}

func (c *FakeProjectcalico) ClusterInformations() internalversion.ClusterInformationInterface {
	return &FakeClusterInformations{c}
}

func (c *FakeProjectcalico) FelixConfigurations() internalversion.FelixConfigurationInterface {
	return &FakeFelixConfigurations{c}
}

func (c *FakeProjectcalico) GlobalNetworkPolicies() internalversion.GlobalNetworkPolicyInterface {
	return &FakeGlobalNetworkPolicies{c}
}

func (c *FakeProjectcalico) GlobalNetworkSets() internalversion.GlobalNetworkSetInterface {
	return &FakeGlobalNetworkSets{c}
}

func (c *FakeProjectcalico) HostEndpoints() internalversion.HostEndpointInterface {
	return &FakeHostEndpoints{c}
}

func (c *FakeProjectcalico) IPPools() internalversion.IPPoolInterface {
	return &FakeIPPools{c}
}

func (c *FakeProjectcalico) KubeControllersConfigurations() internalversion.KubeControllersConfigurationInterface {
	return &FakeKubeControllersConfigurations{c}
}

func (c *FakeProjectcalico) NetworkPolicies(namespace string) internalversion.NetworkPolicyInterface {
	return &FakeNetworkPolicies{c, namespace}
}

func (c *FakeProjectcalico) NetworkSets() internalversion.NetworkSetInterface {
	return &FakeNetworkSets{c}
}

func (c *FakeProjectcalico) Profiles() internalversion.ProfileInterface {
	return &FakeProfiles{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProjectcalico) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
