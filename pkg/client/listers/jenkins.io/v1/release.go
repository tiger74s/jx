// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReleaseLister helps list Releases.
type ReleaseLister interface {
	// List lists all Releases in the indexer.
	List(selector labels.Selector) (ret []*v1.Release, err error)
	// Releases returns an object that can list and get Releases.
	Releases(namespace string) ReleaseNamespaceLister
	ReleaseListerExpansion
}

// releaseLister implements the ReleaseLister interface.
type releaseLister struct {
	indexer cache.Indexer
}

// NewReleaseLister returns a new ReleaseLister.
func NewReleaseLister(indexer cache.Indexer) ReleaseLister {
	return &releaseLister{indexer: indexer}
}

// List lists all Releases in the indexer.
func (s *releaseLister) List(selector labels.Selector) (ret []*v1.Release, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Release))
	})
	return ret, err
}

// Releases returns an object that can list and get Releases.
func (s *releaseLister) Releases(namespace string) ReleaseNamespaceLister {
	return releaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReleaseNamespaceLister helps list and get Releases.
type ReleaseNamespaceLister interface {
	// List lists all Releases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Release, err error)
	// Get retrieves the Release from the indexer for a given namespace and name.
	Get(name string) (*v1.Release, error)
	ReleaseNamespaceListerExpansion
}

// releaseNamespaceLister implements the ReleaseNamespaceLister
// interface.
type releaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Releases in the indexer for a given namespace.
func (s releaseNamespaceLister) List(selector labels.Selector) (ret []*v1.Release, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Release))
	})
	return ret, err
}

// Get retrieves the Release from the indexer for a given namespace and name.
func (s releaseNamespaceLister) Get(name string) (*v1.Release, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("release"), name)
	}
	return obj.(*v1.Release), nil
}
