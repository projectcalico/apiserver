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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
)

// FelixConfigurationLister helps list FelixConfigurations.
// All objects returned here must be treated as read-only.
type FelixConfigurationLister interface {
	// List lists all FelixConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*projectcalico.FelixConfiguration, err error)
	// Get retrieves the FelixConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*projectcalico.FelixConfiguration, error)
	FelixConfigurationListerExpansion
}

// felixConfigurationLister implements the FelixConfigurationLister interface.
type felixConfigurationLister struct {
	indexer cache.Indexer
}

// NewFelixConfigurationLister returns a new FelixConfigurationLister.
func NewFelixConfigurationLister(indexer cache.Indexer) FelixConfigurationLister {
	return &felixConfigurationLister{indexer: indexer}
}

// List lists all FelixConfigurations in the indexer.
func (s *felixConfigurationLister) List(selector labels.Selector) (ret []*projectcalico.FelixConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*projectcalico.FelixConfiguration))
	})
	return ret, err
}

// Get retrieves the FelixConfiguration from the index for a given name.
func (s *felixConfigurationLister) Get(name string) (*projectcalico.FelixConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(projectcalico.Resource("felixconfiguration"), name)
	}
	return obj.(*projectcalico.FelixConfiguration), nil
}
