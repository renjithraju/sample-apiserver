/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	wardle "k8s.io/sample-apiserver/pkg/apis/wardle"
)

// FischerLister helps list Fischers.
type FischerLister interface {
	// List lists all Fischers in the indexer.
	List(selector labels.Selector) (ret []*wardle.Fischer, err error)
	// Get retrieves the Fischer from the index for a given name.
	Get(name string) (*wardle.Fischer, error)
	FischerListerExpansion
}

// fischerLister implements the FischerLister interface.
type fischerLister struct {
	indexer cache.Indexer
}

// NewFischerLister returns a new FischerLister.
func NewFischerLister(indexer cache.Indexer) FischerLister {
	return &fischerLister{indexer: indexer}
}

// List lists all Fischers in the indexer.
func (s *fischerLister) List(selector labels.Selector) (ret []*wardle.Fischer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*wardle.Fischer))
	})
	return ret, err
}

// Get retrieves the Fischer from the index for a given name.
func (s *fischerLister) Get(name string) (*wardle.Fischer, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(wardle.Resource("fischer"), name)
	}
	return obj.(*wardle.Fischer), nil
}
