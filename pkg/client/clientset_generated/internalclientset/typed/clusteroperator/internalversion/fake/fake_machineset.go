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

package fake

import (
	clusteroperator "github.com/openshift/cluster-operator/pkg/apis/clusteroperator"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineSets implements MachineSetInterface
type FakeMachineSets struct {
	Fake *FakeClusteroperator
	ns   string
}

var machinesetsResource = schema.GroupVersionResource{Group: "clusteroperator.openshift.io", Version: "", Resource: "machinesets"}

var machinesetsKind = schema.GroupVersionKind{Group: "clusteroperator.openshift.io", Version: "", Kind: "MachineSet"}

// Get takes name of the machineSet, and returns the corresponding machineSet object, and an error if there is any.
func (c *FakeMachineSets) Get(name string, options v1.GetOptions) (result *clusteroperator.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinesetsResource, c.ns, name), &clusteroperator.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusteroperator.MachineSet), err
}

// List takes label and field selectors, and returns the list of MachineSets that match those selectors.
func (c *FakeMachineSets) List(opts v1.ListOptions) (result *clusteroperator.MachineSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinesetsResource, machinesetsKind, c.ns, opts), &clusteroperator.MachineSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clusteroperator.MachineSetList{}
	for _, item := range obj.(*clusteroperator.MachineSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineSets.
func (c *FakeMachineSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinesetsResource, c.ns, opts))

}

// Create takes the representation of a machineSet and creates it.  Returns the server's representation of the machineSet, and an error, if there is any.
func (c *FakeMachineSets) Create(machineSet *clusteroperator.MachineSet) (result *clusteroperator.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinesetsResource, c.ns, machineSet), &clusteroperator.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusteroperator.MachineSet), err
}

// Update takes the representation of a machineSet and updates it. Returns the server's representation of the machineSet, and an error, if there is any.
func (c *FakeMachineSets) Update(machineSet *clusteroperator.MachineSet) (result *clusteroperator.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinesetsResource, c.ns, machineSet), &clusteroperator.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusteroperator.MachineSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineSets) UpdateStatus(machineSet *clusteroperator.MachineSet) (*clusteroperator.MachineSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinesetsResource, "status", c.ns, machineSet), &clusteroperator.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusteroperator.MachineSet), err
}

// Delete takes name of the machineSet and deletes it. Returns an error if one occurs.
func (c *FakeMachineSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machinesetsResource, c.ns, name), &clusteroperator.MachineSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinesetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &clusteroperator.MachineSetList{})
	return err
}

// Patch applies the patch and returns the patched machineSet.
func (c *FakeMachineSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *clusteroperator.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinesetsResource, c.ns, name, data, subresources...), &clusteroperator.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusteroperator.MachineSet), err
}
