package namespace

import (
	"github.com/wike2019/k8s-wike/config"
	corev1 "k8s.io/api/core/v1"
)

type NsCtl struct {
	*config.Box[*corev1.Namespace]
}

func Ns(box *config.Box[*corev1.Namespace]) *NsCtl {
	return &NsCtl{
		Box: box,
	}
}
func (this *NsCtl) OnAdd(obj interface{}) {
	this.Content.Set(obj.(*corev1.Namespace).Name, obj.(*corev1.Namespace))
}

func (this *NsCtl) OnUpdate(oldObj, newObj interface{}) {
	this.Content.Set(newObj.(*corev1.Namespace).Name, newObj.(*corev1.Namespace))
}
func (this *NsCtl) OnDelete(obj interface{}) {
	this.Content.Delete(obj.(*corev1.Namespace).Name)
}
