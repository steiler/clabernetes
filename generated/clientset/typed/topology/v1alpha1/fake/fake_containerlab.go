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
	"context"

	v1alpha1 "github.com/srl-labs/clabernetes/apis/topology/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContainerlabs implements ContainerlabInterface
type FakeContainerlabs struct {
	Fake *FakeTopologyV1alpha1
	ns   string
}

var containerlabsResource = v1alpha1.SchemeGroupVersion.WithResource("containerlabs")

var containerlabsKind = v1alpha1.SchemeGroupVersion.WithKind("Containerlab")

// Get takes name of the containerlab, and returns the corresponding containerlab object, and an error if there is any.
func (c *FakeContainerlabs) Get(
	ctx context.Context,
	name string,
	options v1.GetOptions,
) (result *v1alpha1.Containerlab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(containerlabsResource, c.ns, name), &v1alpha1.Containerlab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Containerlab), err
}

// List takes label and field selectors, and returns the list of Containerlabs that match those selectors.
func (c *FakeContainerlabs) List(
	ctx context.Context,
	opts v1.ListOptions,
) (result *v1alpha1.ContainerlabList, err error) {
	obj, err := c.Fake.
		Invokes(
			testing.NewListAction(containerlabsResource, containerlabsKind, c.ns, opts),
			&v1alpha1.ContainerlabList{},
		)

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ContainerlabList{ListMeta: obj.(*v1alpha1.ContainerlabList).ListMeta}
	for _, item := range obj.(*v1alpha1.ContainerlabList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerlabs.
func (c *FakeContainerlabs) Watch(
	ctx context.Context,
	opts v1.ListOptions,
) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(containerlabsResource, c.ns, opts))

}

// Create takes the representation of a containerlab and creates it.  Returns the server's representation of the containerlab, and an error, if there is any.
func (c *FakeContainerlabs) Create(
	ctx context.Context,
	containerlab *v1alpha1.Containerlab,
	opts v1.CreateOptions,
) (result *v1alpha1.Containerlab, err error) {
	obj, err := c.Fake.
		Invokes(
			testing.NewCreateAction(containerlabsResource, c.ns, containerlab),
			&v1alpha1.Containerlab{},
		)

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Containerlab), err
}

// Update takes the representation of a containerlab and updates it. Returns the server's representation of the containerlab, and an error, if there is any.
func (c *FakeContainerlabs) Update(
	ctx context.Context,
	containerlab *v1alpha1.Containerlab,
	opts v1.UpdateOptions,
) (result *v1alpha1.Containerlab, err error) {
	obj, err := c.Fake.
		Invokes(
			testing.NewUpdateAction(containerlabsResource, c.ns, containerlab),
			&v1alpha1.Containerlab{},
		)

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Containerlab), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerlabs) UpdateStatus(
	ctx context.Context,
	containerlab *v1alpha1.Containerlab,
	opts v1.UpdateOptions,
) (*v1alpha1.Containerlab, error) {
	obj, err := c.Fake.
		Invokes(
			testing.NewUpdateSubresourceAction(containerlabsResource, "status", c.ns, containerlab),
			&v1alpha1.Containerlab{},
		)

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Containerlab), err
}

// Delete takes name of the containerlab and deletes it. Returns an error if one occurs.
func (c *FakeContainerlabs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(
			testing.NewDeleteActionWithOptions(containerlabsResource, c.ns, name, opts),
			&v1alpha1.Containerlab{},
		)

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerlabs) DeleteCollection(
	ctx context.Context,
	opts v1.DeleteOptions,
	listOpts v1.ListOptions,
) error {
	action := testing.NewDeleteCollectionAction(containerlabsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ContainerlabList{})
	return err
}

// Patch applies the patch and returns the patched containerlab.
func (c *FakeContainerlabs) Patch(
	ctx context.Context,
	name string,
	pt types.PatchType,
	data []byte,
	opts v1.PatchOptions,
	subresources ...string,
) (result *v1alpha1.Containerlab, err error) {
	obj, err := c.Fake.
		Invokes(
			testing.NewPatchSubresourceAction(
				containerlabsResource,
				c.ns,
				name,
				pt,
				data,
				subresources...),
			&v1alpha1.Containerlab{},
		)

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Containerlab), err
}
