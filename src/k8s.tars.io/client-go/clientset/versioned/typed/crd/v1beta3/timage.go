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

package v1beta3

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "k8s.tars.io/client-go/clientset/versioned/scheme"
	v1beta3 "k8s.tars.io/crd/v1beta3"
)

// TImagesGetter has a method to return a TImageInterface.
// A group's client should implement this interface.
type TImagesGetter interface {
	TImages(namespace string) TImageInterface
}

// TImageInterface has methods to work with TImage resources.
type TImageInterface interface {
	Create(ctx context.Context, tImage *v1beta3.TImage, opts v1.CreateOptions) (*v1beta3.TImage, error)
	Update(ctx context.Context, tImage *v1beta3.TImage, opts v1.UpdateOptions) (*v1beta3.TImage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta3.TImage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta3.TImageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TImage, err error)
	TImageExpansion
}

// tImages implements TImageInterface
type tImages struct {
	client rest.Interface
	ns     string
}

// newTImages returns a TImages
func newTImages(c *CrdV1beta3Client, namespace string) *tImages {
	return &tImages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tImage, and returns the corresponding tImage object, and an error if there is any.
func (c *tImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.TImage, err error) {
	result = &v1beta3.TImage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("timages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TImages that match those selectors.
func (c *tImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.TImageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta3.TImageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("timages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tImages.
func (c *tImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("timages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tImage and creates it.  Returns the server's representation of the tImage, and an error, if there is any.
func (c *tImages) Create(ctx context.Context, tImage *v1beta3.TImage, opts v1.CreateOptions) (result *v1beta3.TImage, err error) {
	result = &v1beta3.TImage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("timages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tImage).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tImage and updates it. Returns the server's representation of the tImage, and an error, if there is any.
func (c *tImages) Update(ctx context.Context, tImage *v1beta3.TImage, opts v1.UpdateOptions) (result *v1beta3.TImage, err error) {
	result = &v1beta3.TImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("timages").
		Name(tImage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tImage).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tImage and deletes it. Returns an error if one occurs.
func (c *tImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("timages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("timages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tImage.
func (c *tImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TImage, err error) {
	result = &v1beta3.TImage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("timages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
