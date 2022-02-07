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

package v3

import (
	rest "k8s.io/client-go/rest"

	v3 "github.com/projectcalico/apiserver/pkg/apis/projectcalico/v3"
	"github.com/projectcalico/apiserver/pkg/client/clientset_generated/clientset/scheme"
)

type ProjectcalicoV3Interface interface {
	RESTClient() rest.Interface
	BGPConfigurationsGetter
	BGPPeersGetter
	ClusterInformationsGetter
	FelixConfigurationsGetter
	GlobalNetworkPoliciesGetter
	GlobalNetworkSetsGetter
	HostEndpointsGetter
	IPPoolsGetter
	KubeControllersConfigurationsGetter
	NetworkPoliciesGetter
	NetworkSetsGetter
	ProfilesGetter
}

// ProjectcalicoV3Client is used to interact with features provided by the projectcalico.org group.
type ProjectcalicoV3Client struct {
	restClient rest.Interface
}

func (c *ProjectcalicoV3Client) BGPConfigurations() BGPConfigurationInterface {
	return newBGPConfigurations(c)
}

func (c *ProjectcalicoV3Client) BGPPeers() BGPPeerInterface {
	return newBGPPeers(c)
}

func (c *ProjectcalicoV3Client) ClusterInformations() ClusterInformationInterface {
	return newClusterInformations(c)
}

func (c *ProjectcalicoV3Client) FelixConfigurations() FelixConfigurationInterface {
	return newFelixConfigurations(c)
}

func (c *ProjectcalicoV3Client) GlobalNetworkPolicies() GlobalNetworkPolicyInterface {
	return newGlobalNetworkPolicies(c)
}

func (c *ProjectcalicoV3Client) GlobalNetworkSets() GlobalNetworkSetInterface {
	return newGlobalNetworkSets(c)
}

func (c *ProjectcalicoV3Client) HostEndpoints() HostEndpointInterface {
	return newHostEndpoints(c)
}

func (c *ProjectcalicoV3Client) IPPools() IPPoolInterface {
	return newIPPools(c)
}

func (c *ProjectcalicoV3Client) KubeControllersConfigurations() KubeControllersConfigurationInterface {
	return newKubeControllersConfigurations(c)
}

func (c *ProjectcalicoV3Client) NetworkPolicies(namespace string) NetworkPolicyInterface {
	return newNetworkPolicies(c, namespace)
}

func (c *ProjectcalicoV3Client) NetworkSets(namespace string) NetworkSetInterface {
	return newNetworkSets(c, namespace)
}

func (c *ProjectcalicoV3Client) Profiles() ProfileInterface {
	return newProfiles(c)
}

// NewForConfig creates a new ProjectcalicoV3Client for the given config.
func NewForConfig(c *rest.Config) (*ProjectcalicoV3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ProjectcalicoV3Client{client}, nil
}

// NewForConfigOrDie creates a new ProjectcalicoV3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ProjectcalicoV3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ProjectcalicoV3Client for the given RESTClient.
func New(c rest.Interface) *ProjectcalicoV3Client {
	return &ProjectcalicoV3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v3.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ProjectcalicoV3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
