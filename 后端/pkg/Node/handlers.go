package Node

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
)


//Node相关的handler
type NodeMapHandler struct {
	NodeMap *NodeMapStruct  `inject:"-"`
	NodeService *NodeService `inject:"-"`
}
func(this *NodeMapHandler) OnAdd(obj interface{}){
	this.NodeMap.Add(obj.(*corev1.Node))

	ret:=map[string]interface{}{"result":this.NodeService.ListAllNodes(),"type":"node"}
	wscore.ClientMap.SendAll(ret)

}
func(this *NodeMapHandler) OnUpdate(oldObj, newObj interface{}){
	//重点： 只要update返回true 才会发送 。否则不发送
	if this.NodeMap.Update(newObj.(*corev1.Node)){
		ret:=map[string]interface{}{"result":this.NodeService.ListAllNodes(),"type":"node"}
		wscore.ClientMap.SendAll(ret)
	}
}
func(this *NodeMapHandler) OnDelete(obj interface{}){
	this.NodeMap.Delete(obj.(*corev1.Node))
	ret:=map[string]interface{}{"result":this.NodeService.ListAllNodes(),"type":"node"}
	wscore.ClientMap.SendAll(ret)
}