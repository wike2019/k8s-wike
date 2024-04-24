package informer

import (
	"github.com/wike2019/k8s-wike/src/namespace"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

func InitInformer(client *kubernetes.Clientset, nsctl *namespace.NsCtl) informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(client, 0)
	nsInformer := fact.Core().V1().Namespaces() //监听namespace
	nsInformer.Informer().AddEventHandler(nsctl)
	stopCh := make(chan struct{}, 0)
	go fact.Start(stopCh)
	fact.WaitForCacheSync(stopCh)
	return fact
}
