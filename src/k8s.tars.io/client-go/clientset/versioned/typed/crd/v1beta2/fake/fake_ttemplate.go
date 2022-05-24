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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta2 "k8s.tars.io/crd/v1beta2"
)

// FakeTTemplates implements TTemplateInterface
type FakeTTemplates struct {
	Fake *FakeCrdV1beta2
	ns   string
}

var ttemplatesResource = schema.GroupVersionResource{Group: "crd", Version: "v1beta2", Resource: "ttemplates"}

var ttemplatesKind = schema.GroupVersionKind{Group: "crd", Version: "v1beta2", Kind: "TTemplate"}

// Get takes name of the tTemplate, and returns the corresponding tTemplate object, and an error if there is any.
func (c *FakeTTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.TTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ttemplatesResource, c.ns, name), &v1beta2.TTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.TTemplate), err
}

// List takes label and field selectors, and returns the list of TTemplates that match those selectors.
func (c *FakeTTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.TTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ttemplatesResource, ttemplatesKind, c.ns, opts), &v1beta2.TTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.TTemplateList{ListMeta: obj.(*v1beta2.TTemplateList).ListMeta}
	for _, item := range obj.(*v1beta2.TTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tTemplates.
func (c *FakeTTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ttemplatesResource, c.ns, opts))

}

// Create takes the representation of a tTemplate and creates it.  Returns the server's representation of the tTemplate, and an error, if there is any.
func (c *FakeTTemplates) Create(ctx context.Context, tTemplate *v1beta2.TTemplate, opts v1.CreateOptions) (result *v1beta2.TTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ttemplatesResource, c.ns, tTemplate), &v1beta2.TTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.TTemplate), err
}

// Update takes the representation of a tTemplate and updates it. Returns the server's representation of the tTemplate, and an error, if there is any.
func (c *FakeTTemplates) Update(ctx context.Context, tTemplate *v1beta2.TTemplate, opts v1.UpdateOptions) (result *v1beta2.TTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ttemplatesResource, c.ns, tTemplate), &v1beta2.TTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.TTemplate), err
}

// Delete takes name of the tTemplate and deletes it. Returns an error if one occurs.
func (c *FakeTTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ttemplatesResource, c.ns, name), &v1beta2.TTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ttemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.TTemplateList{})
	return err
}

// Patch applies the patch and returns the patched tTemplate.
func (c *FakeTTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.TTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ttemplatesResource, c.ns, name, pt, data, subresources...), &v1beta2.TTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.TTemplate), err
}
