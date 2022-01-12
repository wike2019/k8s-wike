package Deployment

import "k8sapi/pkg/Pod"

type Deployment struct {
	Name string `json:"name"`
	NameSpace string `json:"name_space"`
	Replicas [3]int32   `json:"replicas"`//3个值，分别是总副本数，可用副本数 ，不可用副本数
	Images string `json:"images"`
	IsComplete bool `json:"is_complete"` //是否完成
	Message string `json:"message"`// 显示错误信息
	CreateTime string `json:"create_time"`
	Pods []*Pod.Pod `json:"pods"`
}