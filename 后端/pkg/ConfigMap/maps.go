package ConfigMap

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/common"
	"k8sapi/pkg/helper"
	"sort"
	"sync"
)

//给configmap的特殊struct
type cm struct {
	cmdata *corev1.ConfigMap
	md5 string  //对cm的data进行md5存储，防止过度更新
}

func newcm(c *corev1.ConfigMap) *cm  {
	helpers:=&helper.Helper{}
	return &cm{
		cmdata:c,//原始对象
		md5:helpers.Md5Data(c.Data),
	}
}
type ConfigMapStruct struct {
	Helper *helper.Helper `inject:"-"`
	data sync.Map         // [ns string] []*cm
}
func(this *ConfigMapStruct) Get(ns string,name string) *corev1.ConfigMap{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*cm){
			if item.cmdata.Name==name{
				return item.cmdata
			}
		}
	}
	return nil
}
func(this *ConfigMapStruct) Add(item *corev1.ConfigMap){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*cm),newcm(item))
		this.data.Store(item.Namespace,list)
		common.AddAutoComplete(fmt.Sprintf("资源类型：%s 命名空间：%s 资源名称 %s","configmap",item.Namespace,item.Name),item.Name,0)
	}else{
		this.data.Store(item.Namespace,[]*cm{newcm(item)})
	}
}
//返回值 是true 或false . true代表有值更新了， 否则返回false
func(this *ConfigMapStruct) Update(item *corev1.ConfigMap) bool {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*cm){
			//这里做判断，如果没变化就不做 更新
			if range_item.cmdata.Name==item.Name && !this.Helper.CmIsEq(range_item.cmdata.Data,item.Data){
				list.([]*cm)[i]=newcm(item)
				return true //代表有值更新了
			}
		}
	}
	return 	  false
}
func(this *ConfigMapStruct) Delete(svc *corev1.ConfigMap){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*cm){
			if range_item.cmdata.Name==svc.Name{
				newList:= append(list.([]*cm)[:i], list.([]*cm)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				common.DeleteAutoComplete(
					fmt.Sprintf("资源类型：%s 命名空间：%s 资源名称 %s",
						"configmap",range_item.cmdata.Namespace,range_item.cmdata.Name))
				break
			}
		}
	}
}
func(this *ConfigMapStruct) ListAll(ns string)[]*ConfigMapModel{
	ret:=[]*ConfigMapModel{}
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*cm)
		sort.Sort(CoreV1ConfigMap(newList))//  按时间倒排序
		for _,cm:=range newList{
			ret=append(ret,&ConfigMapModel{
				Name: cm.cmdata.Name,
				NameSpace: cm.cmdata.Namespace,
				CreateTime: cm.cmdata.CreationTimestamp.Format("2006-01-02 15:04:05"),
				Data: cm.cmdata.Data,
				Labels: cm.cmdata.Labels,
				Annotations: cm.cmdata.Annotations,
			})
		}
	}
	return ret//返回空列表
}

func(this *ConfigMapStruct) GetConfigMap(ns string,name string) ConfigMapModel{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*cm)
		sort.Sort(CoreV1ConfigMap(newList))//  按时间倒排序
		for _,cm:=range newList{
			if cm.cmdata.Name==name{
				return  ConfigMapModel{
					Name: cm.cmdata.Name,
					NameSpace: cm.cmdata.Namespace,
					CreateTime: cm.cmdata.CreationTimestamp.Format("2006-01-02 15:04:05"),
					Data: cm.cmdata.Data,
					Labels: cm.cmdata.Labels,
					Annotations: cm.cmdata.Annotations,
				}
			}

		}
	}
	return  ConfigMapModel{}
}


