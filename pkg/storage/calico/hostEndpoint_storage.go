// Copyright (c) 2019 Tigera, Inc. All rights reserved.

package calico

import (
	"reflect"

	"golang.org/x/net/context"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	etcd "k8s.io/apiserver/pkg/storage/etcd3"
	"k8s.io/apiserver/pkg/storage/storagebackend/factory"

	api "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"github.com/projectcalico/libcalico-go/lib/clientv3"
	"github.com/projectcalico/libcalico-go/lib/options"
	"github.com/projectcalico/libcalico-go/lib/watch"
)

// NewHostEndpointStorage creates a new libcalico-based storage.Interface implementation for HostEndpoints
func NewHostEndpointStorage(opts Options) (registry.DryRunnableStorage, factory.DestroyFunc) {
	c := CreateClientFromConfig()
	createFn := func(ctx context.Context, c clientv3.Interface, obj resourceObject, opts clientOpts) (resourceObject, error) {
		oso := opts.(options.SetOptions)
		res := obj.(*api.HostEndpoint)
		return c.HostEndpoints().Create(ctx, res, oso)
	}
	updateFn := func(ctx context.Context, c clientv3.Interface, obj resourceObject, opts clientOpts) (resourceObject, error) {
		oso := opts.(options.SetOptions)
		res := obj.(*api.HostEndpoint)
		return c.HostEndpoints().Update(ctx, res, oso)
	}
	getFn := func(ctx context.Context, c clientv3.Interface, ns string, name string, opts clientOpts) (resourceObject, error) {
		ogo := opts.(options.GetOptions)
		return c.HostEndpoints().Get(ctx, name, ogo)
	}
	deleteFn := func(ctx context.Context, c clientv3.Interface, ns string, name string, opts clientOpts) (resourceObject, error) {
		odo := opts.(options.DeleteOptions)
		return c.HostEndpoints().Delete(ctx, name, odo)
	}
	listFn := func(ctx context.Context, c clientv3.Interface, opts clientOpts) (resourceListObject, error) {
		olo := opts.(options.ListOptions)
		return c.HostEndpoints().List(ctx, olo)
	}
	watchFn := func(ctx context.Context, c clientv3.Interface, opts clientOpts) (watch.Interface, error) {
		olo := opts.(options.ListOptions)
		return c.HostEndpoints().Watch(ctx, olo)
	}

	dryRunnableStorage := registry.DryRunnableStorage{Storage: &resourceStore{
		client:       c,
		codec:        opts.RESTOptions.StorageConfig.Codec,
		versioner:    etcd.APIObjectVersioner{},
		aapiType:     reflect.TypeOf(api.HostEndpoint{}),
		aapiListType: reflect.TypeOf(api.HostEndpointList{}),
		isNamespaced: false,
		create:       createFn,
		update:       updateFn,
		get:          getFn,
		delete:       deleteFn,
		list:         listFn,
		watch:        watchFn,
		resourceName: "HostEndpoint",
	}, Codec: opts.RESTOptions.StorageConfig.Codec}
	return dryRunnableStorage, func() {}
}
