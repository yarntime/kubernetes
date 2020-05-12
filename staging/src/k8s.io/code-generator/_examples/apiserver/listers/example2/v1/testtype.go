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

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/code-generator/_examples/apiserver/apis/example2/v1"
)

// TestTypeLister helps list TestTypes.
// All objects returned here must be treated as read-only.
type TestTypeLister interface {
	// List lists all TestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.TestType, err error)
	// TestTypes returns an object that can list and get TestTypes.
	TestTypes(namespace string) TestTypeNamespaceLister
	TestTypeListerExpansion
}

// testTypeLister implements the TestTypeLister interface.
type testTypeLister struct {
	indexer cache.Indexer
}

// NewTestTypeLister returns a new TestTypeLister.
func NewTestTypeLister(indexer cache.Indexer) TestTypeLister {
	return &testTypeLister{indexer: indexer}
}

// List lists all TestTypes in the indexer.
func (s *testTypeLister) List(selector labels.Selector) (ret []*v1.TestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TestType))
	})
	return ret, err
}

// TestTypes returns an object that can list and get TestTypes.
func (s *testTypeLister) TestTypes(namespace string) TestTypeNamespaceLister {
	return testTypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TestTypeNamespaceLister helps list and get TestTypes.
// All objects returned here must be treated as read-only.
type TestTypeNamespaceLister interface {
	// List lists all TestTypes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.TestType, err error)
	// Get retrieves the TestType from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.TestType, error)
	TestTypeNamespaceListerExpansion
}

// testTypeNamespaceLister implements the TestTypeNamespaceLister
// interface.
type testTypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TestTypes in the indexer for a given namespace.
func (s testTypeNamespaceLister) List(selector labels.Selector) (ret []*v1.TestType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TestType))
	})
	return ret, err
}

// Get retrieves the TestType from the indexer for a given namespace and name.
func (s testTypeNamespaceLister) Get(name string) (*v1.TestType, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("testtype"), name)
	}
	return obj.(*v1.TestType), nil
}
