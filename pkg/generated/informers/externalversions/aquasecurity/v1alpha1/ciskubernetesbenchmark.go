// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	aquasecurityv1alpha1 "github.com/aquasecurity/starboard/pkg/apis/aquasecurity/v1alpha1"
	versioned "github.com/aquasecurity/starboard/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/aquasecurity/starboard/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/aquasecurity/starboard/pkg/generated/listers/aquasecurity/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CISKubernetesBenchmarkInformer provides access to a shared informer and lister for
// CISKubernetesBenchmarks.
type CISKubernetesBenchmarkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CISKubernetesBenchmarkLister
}

type cISKubernetesBenchmarkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCISKubernetesBenchmarkInformer constructs a new informer for CISKubernetesBenchmark type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCISKubernetesBenchmarkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCISKubernetesBenchmarkInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCISKubernetesBenchmarkInformer constructs a new informer for CISKubernetesBenchmark type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCISKubernetesBenchmarkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AquasecurityV1alpha1().CISKubernetesBenchmarks().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AquasecurityV1alpha1().CISKubernetesBenchmarks().Watch(options)
			},
		},
		&aquasecurityv1alpha1.CISKubernetesBenchmark{},
		resyncPeriod,
		indexers,
	)
}

func (f *cISKubernetesBenchmarkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCISKubernetesBenchmarkInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cISKubernetesBenchmarkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aquasecurityv1alpha1.CISKubernetesBenchmark{}, f.defaultInformer)
}

func (f *cISKubernetesBenchmarkInformer) Lister() v1alpha1.CISKubernetesBenchmarkLister {
	return v1alpha1.NewCISKubernetesBenchmarkLister(f.Informer().GetIndexer())
}
