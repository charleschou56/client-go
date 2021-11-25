/*
Copyright 2021 The KubeVirt Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "kubevirt.io/api/flavor/v1alpha1"
)

// FakeVirtualMachineFlavors implements VirtualMachineFlavorInterface
type FakeVirtualMachineFlavors struct {
	Fake *FakeFlavorV1alpha1
	ns   string
}

var virtualmachineflavorsResource = schema.GroupVersionResource{Group: "flavor.kubevirt.io", Version: "v1alpha1", Resource: "virtualmachineflavors"}

var virtualmachineflavorsKind = schema.GroupVersionKind{Group: "flavor.kubevirt.io", Version: "v1alpha1", Kind: "VirtualMachineFlavor"}

// Get takes name of the virtualMachineFlavor, and returns the corresponding virtualMachineFlavor object, and an error if there is any.
func (c *FakeVirtualMachineFlavors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineFlavor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineflavorsResource, c.ns, name), &v1alpha1.VirtualMachineFlavor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineFlavor), err
}

// List takes label and field selectors, and returns the list of VirtualMachineFlavors that match those selectors.
func (c *FakeVirtualMachineFlavors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineFlavorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineflavorsResource, virtualmachineflavorsKind, c.ns, opts), &v1alpha1.VirtualMachineFlavorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineFlavorList{ListMeta: obj.(*v1alpha1.VirtualMachineFlavorList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineFlavorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineFlavors.
func (c *FakeVirtualMachineFlavors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineflavorsResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineFlavor and creates it.  Returns the server's representation of the virtualMachineFlavor, and an error, if there is any.
func (c *FakeVirtualMachineFlavors) Create(ctx context.Context, virtualMachineFlavor *v1alpha1.VirtualMachineFlavor, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineFlavor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineflavorsResource, c.ns, virtualMachineFlavor), &v1alpha1.VirtualMachineFlavor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineFlavor), err
}

// Update takes the representation of a virtualMachineFlavor and updates it. Returns the server's representation of the virtualMachineFlavor, and an error, if there is any.
func (c *FakeVirtualMachineFlavors) Update(ctx context.Context, virtualMachineFlavor *v1alpha1.VirtualMachineFlavor, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineFlavor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineflavorsResource, c.ns, virtualMachineFlavor), &v1alpha1.VirtualMachineFlavor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineFlavor), err
}

// Delete takes name of the virtualMachineFlavor and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineFlavors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineflavorsResource, c.ns, name), &v1alpha1.VirtualMachineFlavor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineFlavors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineflavorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineFlavorList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineFlavor.
func (c *FakeVirtualMachineFlavors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineFlavor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineflavorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualMachineFlavor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineFlavor), err
}