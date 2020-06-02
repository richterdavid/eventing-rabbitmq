/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	sourcesv1alpha1 "knative.dev/eventing-rabbitmq/source/pkg/apis/sources/v1alpha1"
	versioned "knative.dev/eventing-rabbitmq/source/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing-rabbitmq/source/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "knative.dev/eventing-rabbitmq/source/pkg/client/listers/sources/v1alpha1"
)

// RabbitmqSourceInformer provides access to a shared informer and lister for
// RabbitmqSources.
type RabbitmqSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RabbitmqSourceLister
}

type rabbitmqSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRabbitmqSourceInformer constructs a new informer for RabbitmqSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRabbitmqSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRabbitmqSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRabbitmqSourceInformer constructs a new informer for RabbitmqSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRabbitmqSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().RabbitmqSources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().RabbitmqSources(namespace).Watch(options)
			},
		},
		&sourcesv1alpha1.RabbitmqSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *rabbitmqSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRabbitmqSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rabbitmqSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sourcesv1alpha1.RabbitmqSource{}, f.defaultInformer)
}

func (f *rabbitmqSourceInformer) Lister() v1alpha1.RabbitmqSourceLister {
	return v1alpha1.NewRabbitmqSourceLister(f.Informer().GetIndexer())
}