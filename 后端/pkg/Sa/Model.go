package Sa

import corev1 "k8s.io/api/core/v1"

type Model struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	CreateTime string `json:"create_time"`
	Secrets []corev1.ObjectReference `json:"secrets"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets"`
	Annotations map[string]string `json:"annotations"`
	Labels  map[string]string`json:"labels"`
}
