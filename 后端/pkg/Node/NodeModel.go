package Node
import "k8s.io/api/core/v1"
type NodeUsage struct {
	Pods int64
	Memory float64
	Cpu float64
}

//容量
type NodeCapacity struct {
	Cpu int64
	Memory int64
	Pods int64
}

//保存用
type PostNodeModel struct {
	Name string                    `json:"name"`
	OrginLabels map[string]string `json:"orgin_labels"`
	OrginTaints []v1.Taint `json:"orgin_taints"`
}
type NodeModel struct {
	Name string `json:"name"`
	IP string `json:"ip"`
	HostName string `json:"host_name"`
	CreateTime string `json:"create_time"`
	IsReady string `json:"is_ready"`
	Labels []string `json:"labels"`
	Taints []string `json:"taints"`
	Capacity *NodeCapacity  //容量 包含了cpu 内存和pods数量
	Usage *NodeUsage //资源 使用情况
	OrginLabels map[string]string  //原始标签
	OrginTaints []v1.Taint //原始污点
}
