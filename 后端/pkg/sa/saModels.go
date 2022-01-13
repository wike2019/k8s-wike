package sa

import corev1 "k8s.io/api/core/v1"

type ListModel struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	CreateTime string `json:"createTime"`
	Secrets []corev1.ObjectReference `json:"secrets"`
}
