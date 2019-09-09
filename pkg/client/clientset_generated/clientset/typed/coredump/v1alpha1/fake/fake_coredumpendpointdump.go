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
	coredumpapi "github.com/WanLinghao/api/coredump"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCoredumpEndpointDumps implements CoredumpEndpointDumpInterface
type FakeCoredumpEndpointDumps struct {
	Fake *FakeCoredumpV1alpha1
	ns   string
}

var coredumpendpointdumpsResource = schema.GroupVersionResource{Group: "coredump.fujitsu.com", Version: "v1alpha1", Resource: "coredumpendpointdumps"}

var coredumpendpointdumpsKind = schema.GroupVersionKind{Group: "coredump.fujitsu.com", Version: "v1alpha1", Kind: "CoredumpEndpointDump"}

// Get takes name of the coredumpEndpointDump, and returns the corresponding coredumpEndpointDump object, and an error if there is any.
func (c *FakeCoredumpEndpointDumps) Get(name string, options v1.GetOptions) (result *coredumpapi.CoredumpEndpointDump, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(coredumpendpointdumpsResource, c.ns, name), &coredumpapi.CoredumpEndpointDump{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coredumpapi.CoredumpEndpointDump), err
}

// List takes label and field selectors, and returns the list of CoredumpEndpointDumps that match those selectors.
func (c *FakeCoredumpEndpointDumps) List(opts v1.ListOptions) (result *coredumpapi.CoredumpEndpointDumpList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(coredumpendpointdumpsResource, coredumpendpointdumpsKind, c.ns, opts), &coredumpapi.CoredumpEndpointDumpList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &coredumpapi.CoredumpEndpointDumpList{ListMeta: obj.(*coredumpapi.CoredumpEndpointDumpList).ListMeta}
	for _, item := range obj.(*coredumpapi.CoredumpEndpointDumpList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested coredumpEndpointDumps.
func (c *FakeCoredumpEndpointDumps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(coredumpendpointdumpsResource, c.ns, opts))

}

// Create takes the representation of a coredumpEndpointDump and creates it.  Returns the server's representation of the coredumpEndpointDump, and an error, if there is any.
func (c *FakeCoredumpEndpointDumps) Create(coredumpEndpointDump *coredumpapi.CoredumpEndpointDump) (result *coredumpapi.CoredumpEndpointDump, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(coredumpendpointdumpsResource, c.ns, coredumpEndpointDump), &coredumpapi.CoredumpEndpointDump{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coredumpapi.CoredumpEndpointDump), err
}

// Update takes the representation of a coredumpEndpointDump and updates it. Returns the server's representation of the coredumpEndpointDump, and an error, if there is any.
func (c *FakeCoredumpEndpointDumps) Update(coredumpEndpointDump *coredumpapi.CoredumpEndpointDump) (result *coredumpapi.CoredumpEndpointDump, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(coredumpendpointdumpsResource, c.ns, coredumpEndpointDump), &coredumpapi.CoredumpEndpointDump{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coredumpapi.CoredumpEndpointDump), err
}

// Delete takes name of the coredumpEndpointDump and deletes it. Returns an error if one occurs.
func (c *FakeCoredumpEndpointDumps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(coredumpendpointdumpsResource, c.ns, name), &coredumpapi.CoredumpEndpointDump{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCoredumpEndpointDumps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(coredumpendpointdumpsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &coredumpapi.CoredumpEndpointDumpList{})
	return err
}

// Patch applies the patch and returns the patched coredumpEndpointDump.
func (c *FakeCoredumpEndpointDumps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *coredumpapi.CoredumpEndpointDump, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(coredumpendpointdumpsResource, c.ns, name, data, subresources...), &coredumpapi.CoredumpEndpointDump{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coredumpapi.CoredumpEndpointDump), err
}
