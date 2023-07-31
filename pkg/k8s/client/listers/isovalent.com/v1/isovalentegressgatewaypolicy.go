// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/cilium/cilium/pkg/k8s/apis/isovalent.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IsovalentEgressGatewayPolicyLister helps list IsovalentEgressGatewayPolicies.
// All objects returned here must be treated as read-only.
type IsovalentEgressGatewayPolicyLister interface {
	// List lists all IsovalentEgressGatewayPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.IsovalentEgressGatewayPolicy, err error)
	// Get retrieves the IsovalentEgressGatewayPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.IsovalentEgressGatewayPolicy, error)
	IsovalentEgressGatewayPolicyListerExpansion
}

// isovalentEgressGatewayPolicyLister implements the IsovalentEgressGatewayPolicyLister interface.
type isovalentEgressGatewayPolicyLister struct {
	indexer cache.Indexer
}

// NewIsovalentEgressGatewayPolicyLister returns a new IsovalentEgressGatewayPolicyLister.
func NewIsovalentEgressGatewayPolicyLister(indexer cache.Indexer) IsovalentEgressGatewayPolicyLister {
	return &isovalentEgressGatewayPolicyLister{indexer: indexer}
}

// List lists all IsovalentEgressGatewayPolicies in the indexer.
func (s *isovalentEgressGatewayPolicyLister) List(selector labels.Selector) (ret []*v1.IsovalentEgressGatewayPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.IsovalentEgressGatewayPolicy))
	})
	return ret, err
}

// Get retrieves the IsovalentEgressGatewayPolicy from the index for a given name.
func (s *isovalentEgressGatewayPolicyLister) Get(name string) (*v1.IsovalentEgressGatewayPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("isovalentegressgatewaypolicy"), name)
	}
	return obj.(*v1.IsovalentEgressGatewayPolicy), nil
}
