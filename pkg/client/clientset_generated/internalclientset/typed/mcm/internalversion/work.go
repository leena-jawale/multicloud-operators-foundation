// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	mcm "github.com/open-cluster-management/multicloud-operators-foundation/pkg/apis/mcm"
	scheme "github.com/open-cluster-management/multicloud-operators-foundation/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorksGetter has a method to return a WorkInterface.
// A group's client should implement this interface.
type WorksGetter interface {
	Works(namespace string) WorkInterface
}

// WorkInterface has methods to work with Work resources.
type WorkInterface interface {
	Create(*mcm.Work) (*mcm.Work, error)
	Update(*mcm.Work) (*mcm.Work, error)
	UpdateStatus(*mcm.Work) (*mcm.Work, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*mcm.Work, error)
	List(opts v1.ListOptions) (*mcm.WorkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mcm.Work, err error)
	WorkExpansion
}

// works implements WorkInterface
type works struct {
	client rest.Interface
	ns     string
}

// newWorks returns a Works
func newWorks(c *McmClient, namespace string) *works {
	return &works{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the work, and returns the corresponding work object, and an error if there is any.
func (c *works) Get(name string, options v1.GetOptions) (result *mcm.Work, err error) {
	result = &mcm.Work{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("works").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Works that match those selectors.
func (c *works) List(opts v1.ListOptions) (result *mcm.WorkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &mcm.WorkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("works").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested works.
func (c *works) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("works").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a work and creates it.  Returns the server's representation of the work, and an error, if there is any.
func (c *works) Create(work *mcm.Work) (result *mcm.Work, err error) {
	result = &mcm.Work{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("works").
		Body(work).
		Do().
		Into(result)
	return
}

// Update takes the representation of a work and updates it. Returns the server's representation of the work, and an error, if there is any.
func (c *works) Update(work *mcm.Work) (result *mcm.Work, err error) {
	result = &mcm.Work{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("works").
		Name(work.Name).
		Body(work).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *works) UpdateStatus(work *mcm.Work) (result *mcm.Work, err error) {
	result = &mcm.Work{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("works").
		Name(work.Name).
		SubResource("status").
		Body(work).
		Do().
		Into(result)
	return
}

// Delete takes name of the work and deletes it. Returns an error if one occurs.
func (c *works) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("works").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *works) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("works").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched work.
func (c *works) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mcm.Work, err error) {
	result = &mcm.Work{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("works").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
