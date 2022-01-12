package Pod

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"reflect"
	"sort"
	"sync"
)






// 保存Pod集合
type PodMapStruct struct {
	data sync.Map  // [key string] []*v1.Pod    key=>namespace
}
func(this *PodMapStruct) ListByNs(ns string) []*corev1.Pod{
	if list,ok:=this.data.Load(ns);ok{  //取数据
		ret:=list.([]*corev1.Pod) //断言
		sort.Sort(CoreV1Pods(ret))//排序
		return ret
	}
	return nil
}
func(this *PodMapStruct) Get(ns string,podName string) *corev1.Pod{
	if list,ok:=this.data.Load(ns);ok{  //根据命名空间和pod名称获取pod对象
		for _,pod:=range list.([]*corev1.Pod){
			if pod.Name==podName{
				return pod
			}
		}
	}
	return nil
}
func(this *PodMapStruct) Add(pod *corev1.Pod){
	if list,ok:=this.data.Load(pod.Namespace);ok{
		list=append(list.([]*corev1.Pod),pod) //取出列表添加再存回去
		this.data.Store(pod.Namespace,list)
	}else{
		this.data.Store(pod.Namespace,[]*corev1.Pod{pod}) //第一次添加
	}
}
func(this *PodMapStruct) Update(pod *corev1.Pod) error {
	if list,ok:=this.data.Load(pod.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.Pod){ //更新
			if range_pod.Name==pod.Name{
				list.([]*corev1.Pod)[i]=pod
			}
		}
		return nil
	}
	return fmt.Errorf("Pod-%s not found",pod.Name)
}
func(this *PodMapStruct) Delete(pod *corev1.Pod){
	if list,ok:=this.data.Load(pod.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.Pod){
			if range_pod.Name==pod.Name{
				newList:= append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...) //删除
				this.data.Store(pod.Namespace,newList)
				break
			}
		}
	}
}

//根据节点名称 获取pods数量
func(this *PodMapStruct) GetNum(nodeName string) (num int ){
	this.data.Range(func(key, value interface{}) bool {
		list:=value.([]*corev1.Pod)
		for _,pod:=range list{
			if pod.Spec.NodeName==nodeName{
				num++
			}
		}
		return true
	})
	return
}


//根据标签获取 POD列表
func(this *PodMapStruct) ListByLabels(ns string,labels []map[string]string) ([]*corev1.Pod,error){
	ret:=make([]*corev1.Pod,0)
	if list,ok:=this.data.Load(ns);ok {
		for _,pod:=range list.([]*corev1.Pod){
			for _,label:=range labels{
				if reflect.DeepEqual(pod.Labels,label){  //标签完全匹配
					ret=append(ret,pod)
				}
			}
		}
		return ret,nil
	}

	return nil,fmt.Errorf("pods not found ")
}


