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

package internalversion

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
	internalclientset "github.com/projectcalico/apiserver/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/projectcalico/apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/projectcalico/apiserver/pkg/client/listers_generated/projectcalico/internalversion"
)

// IPPoolInformer provides access to a shared informer and lister for
// IPPools.
type IPPoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.IPPoolLister
}

type iPPoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIPPoolInformer constructs a new informer for IPPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIPPoolInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIPPoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIPPoolInformer constructs a new informer for IPPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIPPoolInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Projectcalico().IPPools().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Projectcalico().IPPools().Watch(context.TODO(), options)
			},
		},
		&projectcalico.IPPool{},
		resyncPeriod,
		indexers,
	)
}

func (f *iPPoolInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIPPoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iPPoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcalico.IPPool{}, f.defaultInformer)
}

func (f *iPPoolInformer) Lister() internalversion.IPPoolLister {
	return internalversion.NewIPPoolLister(f.Informer().GetIndexer())
}
