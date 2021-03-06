/*
Copyright 2018 Openstorage.org

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

package v1alpha1

import (
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DataExportLister helps list DataExports.
// All objects returned here must be treated as read-only.
type DataExportLister interface {
	// List lists all DataExports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataExport, err error)
	// DataExports returns an object that can list and get DataExports.
	DataExports(namespace string) DataExportNamespaceLister
	DataExportListerExpansion
}

// dataExportLister implements the DataExportLister interface.
type dataExportLister struct {
	indexer cache.Indexer
}

// NewDataExportLister returns a new DataExportLister.
func NewDataExportLister(indexer cache.Indexer) DataExportLister {
	return &dataExportLister{indexer: indexer}
}

// List lists all DataExports in the indexer.
func (s *dataExportLister) List(selector labels.Selector) (ret []*v1alpha1.DataExport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataExport))
	})
	return ret, err
}

// DataExports returns an object that can list and get DataExports.
func (s *dataExportLister) DataExports(namespace string) DataExportNamespaceLister {
	return dataExportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataExportNamespaceLister helps list and get DataExports.
// All objects returned here must be treated as read-only.
type DataExportNamespaceLister interface {
	// List lists all DataExports in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataExport, err error)
	// Get retrieves the DataExport from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DataExport, error)
	DataExportNamespaceListerExpansion
}

// dataExportNamespaceLister implements the DataExportNamespaceLister
// interface.
type dataExportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataExports in the indexer for a given namespace.
func (s dataExportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataExport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataExport))
	})
	return ret, err
}

// Get retrieves the DataExport from the indexer for a given namespace and name.
func (s dataExportNamespaceLister) Get(name string) (*v1alpha1.DataExport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dataexport"), name)
	}
	return obj.(*v1alpha1.DataExport), nil
}
