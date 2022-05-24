/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "k8s.tars.io/crd/v1beta1"
)

// TConfigLister helps list TConfigs.
type TConfigLister interface {
	// List lists all TConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.TConfig, err error)
	// TConfigs returns an object that can list and get TConfigs.
	TConfigs(namespace string) TConfigNamespaceLister
	TConfigListerExpansion
}

// tConfigLister implements the TConfigLister interface.
type tConfigLister struct {
	indexer cache.Indexer
}

// NewTConfigLister returns a new TConfigLister.
func NewTConfigLister(indexer cache.Indexer) TConfigLister {
	return &tConfigLister{indexer: indexer}
}

// List lists all TConfigs in the indexer.
func (s *tConfigLister) List(selector labels.Selector) (ret []*v1beta1.TConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TConfig))
	})
	return ret, err
}

// TConfigs returns an object that can list and get TConfigs.
func (s *tConfigLister) TConfigs(namespace string) TConfigNamespaceLister {
	return tConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TConfigNamespaceLister helps list and get TConfigs.
type TConfigNamespaceLister interface {
	// List lists all TConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.TConfig, err error)
	// Get retrieves the TConfig from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.TConfig, error)
	TConfigNamespaceListerExpansion
}

// tConfigNamespaceLister implements the TConfigNamespaceLister
// interface.
type tConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TConfigs in the indexer for a given namespace.
func (s tConfigNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.TConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TConfig))
	})
	return ret, err
}

// Get retrieves the TConfig from the indexer for a given namespace and name.
func (s tConfigNamespaceLister) Get(name string) (*v1beta1.TConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("tconfig"), name)
	}
	return obj.(*v1beta1.TConfig), nil
}
