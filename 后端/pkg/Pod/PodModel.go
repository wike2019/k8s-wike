package Pod

import corev1 "k8s.io/api/core/v1"

type Pod struct {
	Name string `json:"name"`
	NameSpace string `json:"name_space"` //新增一个命名空间
	Images string `json:"images"`
	NodeName string  `json:"node_name"`
	IP []string `json:"ip"`// 第一个是 POD IP 第二个是 node ip
	Phase string  `json:"phase"`// pod 当前所处的阶段
	IsReady bool `json:"is_ready"`//判断pod 是否就绪
	Message []*corev1.Event`json:"message"`
	CreateTime string `json:"create_time"`
}