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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/WanLinghao/fujitsu-coredump/pkg/apis/coredump/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCoredumpEndpoints implements CoredumpEndpointInterface
type FakeCoredumpEndpoints struct {
	Fake *FakeCoredumpV1alpha1
	ns   string
}

var coredumpendpointsResource = schema.GroupVersionResource{Group: "coredump.fujitsu.com", Version: "v1alpha1", Resource: "coredumpendpoints"}

var coredumpendpointsKind = schema.GroupVersionKind{Group: "coredump.fujitsu.com", Version: "v1alpha1", Kind: "CoredumpEndpoint"}

// Get takes name of the coredumpEndpoint, and returns the corresponding coredumpEndpoint object, and an error if there is any.
func (c *FakeCoredumpEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.CoredumpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(coredumpendpointsResource, c.ns, name), &v1alpha1.CoredumpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CoredumpEndpoint), err
}

// List takes label and field selectors, and returns the list of CoredumpEndpoints that match those selectors.
func (c *FakeCoredumpEndpoints) List(opts v1.ListOptions) (result *v1alpha1.CoredumpEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(coredumpendpointsResource, coredumpendpointsKind, c.ns, opts), &v1alpha1.CoredumpEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CoredumpEndpointList{ListMeta: obj.(*v1alpha1.CoredumpEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.CoredumpEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested coredumpEndpoints.
func (c *FakeCoredumpEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(coredumpendpointsResource, c.ns, opts))

}

// Create takes the representation of a coredumpEndpoint and creates it.  Returns the server's representation of the coredumpEndpoint, and an error, if there is any.
func (c *FakeCoredumpEndpoints) Create(coredumpEndpoint *v1alpha1.CoredumpEndpoint) (result *v1alpha1.CoredumpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(coredumpendpointsResource, c.ns, coredumpEndpoint), &v1alpha1.CoredumpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CoredumpEndpoint), err
}

// Update takes the representation of a coredumpEndpoint and updates it. Returns the server's representation of the coredumpEndpoint, and an error, if there is any.
func (c *FakeCoredumpEndpoints) Update(coredumpEndpoint *v1alpha1.CoredumpEndpoint) (result *v1alpha1.CoredumpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(coredumpendpointsResource, c.ns, coredumpEndpoint), &v1alpha1.CoredumpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CoredumpEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCoredumpEndpoints) UpdateStatus(coredumpEndpoint *v1alpha1.CoredumpEndpoint) (*v1alpha1.CoredumpEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(coredumpendpointsResource, "status", c.ns, coredumpEndpoint), &v1alpha1.CoredumpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CoredumpEndpoint), err
}

// Delete takes name of the coredumpEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeCoredumpEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(coredumpendpointsResource, c.ns, name), &v1alpha1.CoredumpEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCoredumpEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(coredumpendpointsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CoredumpEndpointList{})
	return err
}

// Patch applies the patch and returns the patched coredumpEndpoint.
func (c *FakeCoredumpEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CoredumpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(coredumpendpointsResource, c.ns, name, data, subresources...), &v1alpha1.CoredumpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CoredumpEndpoint), err
}
