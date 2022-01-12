package Node

import (
	"k8s.io/api/core/v1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"k8sapi/pkg/Common"

)

//@service
type NodeService struct {
	NodeMap *NodeMapStruct `inject:"-"`
	Helper *Common.Helper `inject:"-"`//帮助函数 用于分页
	NodeHelper *Helper `inject:"-"`//帮助函数 用于分页
	Metric *versioned.Clientset `inject:"-"`
}

func NewNodeService() *NodeService {
	return &NodeService{}
}
//显示所有节点
func(this *NodeService) ListAllNodes() []*NodeModel{
    list:=this.NodeMap.ListAll()

    ret:=make([]*NodeModel,len(list))
    for i,node:=range list{
    	ret[i]=&NodeModel{
    		Name:node.Name,
			IP:node.Status.Addresses[0].Address,
			HostName:node.Status.Addresses[1].Address,
			Labels: this.NodeHelper.FilterLables(node.Labels),
			Taints:this.NodeHelper.FilterTaints(node.Spec.Taints),
			CreateTime:node.CreationTimestamp.Format("2006-01-02 15:04:05"),
			IsReady:string(node.Status.Conditions[len(node.Status.Conditions)-1].Type),
			Capacity:&NodeCapacity{
    			Cpu:node.Status.Capacity.Cpu().Value(),
    			Memory: node.Status.Capacity.Memory().Value(),
    			Pods: node.Status.Capacity.Pods().Value(),
			},

			Usage:this.NodeHelper.GetNodeUsage(this.Metric,node),
		}
	}
    return ret
}

//加载node信息， 给编辑用的
func(this *NodeService) LoadNode(nodeName string ) *NodeModel{
	node:= this.NodeMap.Get(nodeName)
	return &NodeModel{
		Name:node.Name,
		IP:node.Status.Addresses[0].Address,
		HostName:node.Status.Addresses[1].Address,
		OrginLabels:node.Labels,
		OrginTaints:node.Spec.Taints,
		Capacity:&NodeCapacity{
			Cpu:node.Status.Capacity.Cpu().Value(),
			Memory: node.Status.Capacity.Memory().Value(),
			Pods: node.Status.Capacity.Pods().Value(),
		},
		Usage:this.NodeHelper.GetNodeUsage(this.Metric,node),
		CreateTime:node.CreationTimestamp.Format("2006-01-02 15:04:05"),
	}
}

//保存时用的
func(this *NodeService) LoadOriginNode(nodeName string ) *v1.Node{
	return this.NodeMap.Get(nodeName)
}