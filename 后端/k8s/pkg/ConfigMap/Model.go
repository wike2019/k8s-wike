package ConfigMap


//列表用
type ConfigMapModel struct {
	Name string `json:"name"`
	NameSpace string `json:"namespace"`
	CreateTime string       `json:"create_time"`
	Data map[string]string `json:"data"`
	Annotations map[string]string `json:"annotations"`
	Labels  map[string]string`json:"labels"`
}