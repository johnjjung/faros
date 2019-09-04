/*
Copyright 2018 Pusher Ltd.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/pusher/faros/pkg/apis/faros/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterGitTrackLister helps list ClusterGitTracks.
type ClusterGitTrackLister interface {
	// List lists all ClusterGitTracks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterGitTrack, err error)
	// Get retrieves the ClusterGitTrack from the index for a given name.
	Get(name string) (*v1alpha1.ClusterGitTrack, error)
	ClusterGitTrackListerExpansion
}

// clusterGitTrackLister implements the ClusterGitTrackLister interface.
type clusterGitTrackLister struct {
	indexer cache.Indexer
}

// NewClusterGitTrackLister returns a new ClusterGitTrackLister.
func NewClusterGitTrackLister(indexer cache.Indexer) ClusterGitTrackLister {
	return &clusterGitTrackLister{indexer: indexer}
}

// List lists all ClusterGitTracks in the indexer.
func (s *clusterGitTrackLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterGitTrack, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterGitTrack))
	})
	return ret, err
}

// Get retrieves the ClusterGitTrack from the index for a given name.
func (s *clusterGitTrackLister) Get(name string) (*v1alpha1.ClusterGitTrack, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustergittrack"), name)
	}
	return obj.(*v1alpha1.ClusterGitTrack), nil
}