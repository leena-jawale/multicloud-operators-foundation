// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/open-cluster-management/multicloud-operators-foundation/pkg/client/clientset_generated/clientset/typed/mcm/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMcmV1beta1 struct {
	*testing.Fake
}

func (c *FakeMcmV1beta1) ClusterJoinRequests() v1beta1.ClusterJoinRequestInterface {
	return &FakeClusterJoinRequests{c}
}

func (c *FakeMcmV1beta1) ClusterStatuses(namespace string) v1beta1.ClusterStatusInterface {
	return &FakeClusterStatuses{c, namespace}
}

func (c *FakeMcmV1beta1) ResourceViews(namespace string) v1beta1.ResourceViewInterface {
	return &FakeResourceViews{c, namespace}
}

func (c *FakeMcmV1beta1) Works(namespace string) v1beta1.WorkInterface {
	return &FakeWorks{c, namespace}
}

func (c *FakeMcmV1beta1) WorkSets(namespace string) v1beta1.WorkSetInterface {
	return &FakeWorkSets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMcmV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
