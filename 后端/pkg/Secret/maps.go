package Secret

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/helper"
	"sort"
	"sync"
)

//SecretMap
type MapStruct struct {
	data sync.Map         // [ns string] []*v1.Secret
	Helper *helper.Helper `inject:"-"`
	SecretHelper *Helper  `inject:"-"`
}
func(this *MapStruct) Get(ns string,name string) *corev1.Secret{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*corev1.Secret){
			if item.Name==name{
				return item
			}
		}
	}
	return nil
}
func(this *MapStruct) Add(item *corev1.Secret){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*corev1.Secret),item)
		this.data.Store(item.Namespace,list)
	}else{
		this.data.Store(item.Namespace,[]*corev1.Secret{item})
	}
}
func(this *MapStruct) Update(item *corev1.Secret) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*corev1.Secret){
			if range_item.Name==item.Name{
				list.([]*corev1.Secret)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("Secret-%s not found",item.Name)
}
func(this *MapStruct) Delete(svc *corev1.Secret){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*corev1.Secret){
			if range_item.Name==svc.Name{
				newList:= append(list.([]*corev1.Secret)[:i], list.([]*corev1.Secret)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *MapStruct) ListAll(ns string)[]*SecretModel{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*corev1.Secret)
		sort.Sort(CoreV1Secret(newList))//  按时间倒排序
		ret:=make([]*SecretModel,len(newList))
		for i,item:=range newList{
			ret[i]=&SecretModel{
				Name:item.Name,
				CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
				NameSpace:item.Namespace,
				Type: SECRET_TYPE[string(item.Type)],
				Data: item.Data,
				StringData:item.StringData,
				TlsData: this.SecretHelper.ParseCert(string(item.Type),item.Data["tls.crt"]),
				Labels:item.Labels,
				Annotations:item.Annotations,
			}
		}
		return ret
	}
	return []*SecretModel{} //返回空列表
}

func(this *MapStruct) GetSecret(ns string,name string) SecretModel{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*corev1.Secret)
		for _,item:=range newList{
			if item.Name==name{
				return  SecretModel{
					Name:item.Name,
					CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
					NameSpace:item.Namespace,
					Type: SECRET_TYPE[string(item.Type)],
					Data: item.Data,
					StringData:item.StringData,
					TlsData: this.SecretHelper.ParseCert(string(item.Type),item.Data["tls.crt"]),
					TypeRaw: string(item.Type),
					Labels:item.Labels,
					Annotations:item.Annotations,
				}

			}
		}
	}
	return SecretModel{} //返回空列表
}


