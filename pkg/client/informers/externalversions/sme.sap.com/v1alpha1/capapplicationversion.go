/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	smesapcomv1alpha1 "github.com/sap/cap-operator/pkg/apis/sme.sap.com/v1alpha1"
	versioned "github.com/sap/cap-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/sap/cap-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/sap/cap-operator/pkg/client/listers/sme.sap.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CAPApplicationVersionInformer provides access to a shared informer and lister for
// CAPApplicationVersions.
type CAPApplicationVersionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CAPApplicationVersionLister
}

type cAPApplicationVersionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCAPApplicationVersionInformer constructs a new informer for CAPApplicationVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCAPApplicationVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCAPApplicationVersionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCAPApplicationVersionInformer constructs a new informer for CAPApplicationVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCAPApplicationVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SmeV1alpha1().CAPApplicationVersions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SmeV1alpha1().CAPApplicationVersions(namespace).Watch(context.TODO(), options)
			},
		},
		&smesapcomv1alpha1.CAPApplicationVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *cAPApplicationVersionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCAPApplicationVersionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cAPApplicationVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&smesapcomv1alpha1.CAPApplicationVersion{}, f.defaultInformer)
}

func (f *cAPApplicationVersionInformer) Lister() v1alpha1.CAPApplicationVersionLister {
	return v1alpha1.NewCAPApplicationVersionLister(f.Informer().GetIndexer())
}