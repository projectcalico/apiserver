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

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v3 "github.com/projectcalico/apiserver/pkg/apis/projectcalico/v3"
)

// ClusterInformationLister helps list ClusterInformations.
// All objects returned here must be treated as read-only.
type ClusterInformationLister interface {
	// List lists all ClusterInformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.ClusterInformation, err error)
	// Get retrieves the ClusterInformation from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.ClusterInformation, error)
	ClusterInformationListerExpansion
}

// clusterInformationLister implements the ClusterInformationLister interface.
type clusterInformationLister struct {
	indexer cache.Indexer
}

// NewClusterInformationLister returns a new ClusterInformationLister.
func NewClusterInformationLister(indexer cache.Indexer) ClusterInformationLister {
	return &clusterInformationLister{indexer: indexer}
}

// List lists all ClusterInformations in the indexer.
func (s *clusterInformationLister) List(selector labels.Selector) (ret []*v3.ClusterInformation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ClusterInformation))
	})
	return ret, err
}

// Get retrieves the ClusterInformation from the index for a given name.
func (s *clusterInformationLister) Get(name string) (*v3.ClusterInformation, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("clusterinformation"), name)
	}
	return obj.(*v3.ClusterInformation), nil
}
