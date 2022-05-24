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
	v1beta3 "k8s.tars.io/crd/v1beta3"
)

// FakeTAccounts implements TAccountInterface
type FakeTAccounts struct {
	Fake *FakeCrdV1beta3
	ns   string
}

var taccountsResource = schema.GroupVersionResource{Group: "crd", Version: "v1beta3", Resource: "taccounts"}

var taccountsKind = schema.GroupVersionKind{Group: "crd", Version: "v1beta3", Kind: "TAccount"}

// Get takes name of the tAccount, and returns the corresponding tAccount object, and an error if there is any.
func (c *FakeTAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.TAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(taccountsResource, c.ns, name), &v1beta3.TAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TAccount), err
}

// List takes label and field selectors, and returns the list of TAccounts that match those selectors.
func (c *FakeTAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.TAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(taccountsResource, taccountsKind, c.ns, opts), &v1beta3.TAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta3.TAccountList{ListMeta: obj.(*v1beta3.TAccountList).ListMeta}
	for _, item := range obj.(*v1beta3.TAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tAccounts.
func (c *FakeTAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(taccountsResource, c.ns, opts))

}

// Create takes the representation of a tAccount and creates it.  Returns the server's representation of the tAccount, and an error, if there is any.
func (c *FakeTAccounts) Create(ctx context.Context, tAccount *v1beta3.TAccount, opts v1.CreateOptions) (result *v1beta3.TAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(taccountsResource, c.ns, tAccount), &v1beta3.TAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TAccount), err
}

// Update takes the representation of a tAccount and updates it. Returns the server's representation of the tAccount, and an error, if there is any.
func (c *FakeTAccounts) Update(ctx context.Context, tAccount *v1beta3.TAccount, opts v1.UpdateOptions) (result *v1beta3.TAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(taccountsResource, c.ns, tAccount), &v1beta3.TAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TAccount), err
}

// Delete takes name of the tAccount and deletes it. Returns an error if one occurs.
func (c *FakeTAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(taccountsResource, c.ns, name), &v1beta3.TAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(taccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta3.TAccountList{})
	return err
}

// Patch applies the patch and returns the patched tAccount.
func (c *FakeTAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(taccountsResource, c.ns, name, pt, data, subresources...), &v1beta3.TAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TAccount), err
}
