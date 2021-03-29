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

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	projectcalicov3 "github.com/projectcalico/apiserver/pkg/apis/projectcalico/v3"
	clientset "github.com/projectcalico/apiserver/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/projectcalico/apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
	v3 "github.com/projectcalico/apiserver/pkg/client/listers_generated/projectcalico/v3"
)

// NetworkSetInformer provides access to a shared informer and lister for
// NetworkSets.
type NetworkSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v3.NetworkSetLister
}

type networkSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkSetInformer constructs a new informer for NetworkSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkSetInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkSetInformer constructs a new informer for NetworkSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkSetInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().NetworkSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().NetworkSets(namespace).Watch(context.TODO(), options)
			},
		},
		&projectcalicov3.NetworkSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkSetInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcalicov3.NetworkSet{}, f.defaultInformer)
}

func (f *networkSetInformer) Lister() v3.NetworkSetLister {
	return v3.NewNetworkSetLister(f.Informer().GetIndexer())
}
