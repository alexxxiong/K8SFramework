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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta3

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.tars.io/client-go/clientset/versioned"
	internalinterfaces "k8s.tars.io/client-go/informers/externalversions/internalinterfaces"
	v1beta3 "k8s.tars.io/client-go/listers/crd/v1beta3"
	crdv1beta3 "k8s.tars.io/crd/v1beta3"
)

// TImageInformer provides access to a shared informer and lister for
// TImages.
type TImageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta3.TImageLister
}

type tImageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTImageInformer constructs a new informer for TImage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTImageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTImageInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTImageInformer constructs a new informer for TImage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTImageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1beta3().TImages(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1beta3().TImages(namespace).Watch(context.TODO(), options)
			},
		},
		&crdv1beta3.TImage{},
		resyncPeriod,
		indexers,
	)
}

func (f *tImageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTImageInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tImageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crdv1beta3.TImage{}, f.defaultInformer)
}

func (f *tImageInformer) Lister() v1beta3.TImageLister {
	return v1beta3.NewTImageLister(f.Informer().GetIndexer())
}
