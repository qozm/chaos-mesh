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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/cwen0/chaos-operator/pkg/apis/pingcap.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodChaoses implements PodChaosInterface
type FakePodChaoses struct {
	Fake *FakePingcapV1alpha1
	ns   string
}

var podchaosesResource = schema.GroupVersionResource{Group: "pingcap.com", Version: "v1alpha1", Resource: "podchaoses"}

var podchaosesKind = schema.GroupVersionKind{Group: "pingcap.com", Version: "v1alpha1", Kind: "PodChaos"}

// Get takes name of the podChaos, and returns the corresponding podChaos object, and an error if there is any.
func (c *FakePodChaoses) Get(name string, options v1.GetOptions) (result *v1alpha1.PodChaos, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podchaosesResource, c.ns, name), &v1alpha1.PodChaos{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaos), err
}

// List takes label and field selectors, and returns the list of PodChaoses that match those selectors.
func (c *FakePodChaoses) List(opts v1.ListOptions) (result *v1alpha1.PodChaosList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podchaosesResource, podchaosesKind, c.ns, opts), &v1alpha1.PodChaosList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodChaosList{ListMeta: obj.(*v1alpha1.PodChaosList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodChaosList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podChaoses.
func (c *FakePodChaoses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podchaosesResource, c.ns, opts))

}

// Create takes the representation of a podChaos and creates it.  Returns the server's representation of the podChaos, and an error, if there is any.
func (c *FakePodChaoses) Create(podChaos *v1alpha1.PodChaos) (result *v1alpha1.PodChaos, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podchaosesResource, c.ns, podChaos), &v1alpha1.PodChaos{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaos), err
}

// Update takes the representation of a podChaos and updates it. Returns the server's representation of the podChaos, and an error, if there is any.
func (c *FakePodChaoses) Update(podChaos *v1alpha1.PodChaos) (result *v1alpha1.PodChaos, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podchaosesResource, c.ns, podChaos), &v1alpha1.PodChaos{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaos), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodChaoses) UpdateStatus(podChaos *v1alpha1.PodChaos) (*v1alpha1.PodChaos, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podchaosesResource, "status", c.ns, podChaos), &v1alpha1.PodChaos{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaos), err
}

// Delete takes name of the podChaos and deletes it. Returns an error if one occurs.
func (c *FakePodChaoses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podchaosesResource, c.ns, name), &v1alpha1.PodChaos{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodChaoses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podchaosesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodChaosList{})
	return err
}

// Patch applies the patch and returns the patched podChaos.
func (c *FakePodChaoses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodChaos, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podchaosesResource, c.ns, name, data, subresources...), &v1alpha1.PodChaos{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodChaos), err
}
