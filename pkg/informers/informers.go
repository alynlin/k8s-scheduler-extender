package informers

import (
	k8s "github.com/k8s-scheduler-extender/pkg/simple/client"
	"k8s.io/client-go/informers"
	"sync"
	"time"
)

const defaultResync = 30 * time.Second

var (
	k8sOnce         sync.Once
	informerFactory informers.SharedInformerFactory
)

func SharedInformerFactory() informers.SharedInformerFactory {
	k8sOnce.Do(func() {
		k8sClient := k8s.Client()
		informerFactory = informers.NewSharedInformerFactory(k8sClient, defaultResync)
	})
	return informerFactory
}
