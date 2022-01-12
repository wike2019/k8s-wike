package Other

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"sync"
)


var mutex sync.Mutex
var AutoCompleteData =make([]interface{},0)

type OtherCtl struct{

}
func NewOtherCtl() *OtherCtl {
  return &OtherCtl{
  }
}
func AddAutoComplete(meta string,key string,score int){
	mutex.Lock()
	defer mutex.Unlock()
	data:=map[string]interface{}{"meta":meta,"caption":key,"value":key,"score":score}
	AutoCompleteData=append(AutoCompleteData,data)
}
func AddSySAutoComplete(key string){
	mutex.Lock()
	defer mutex.Unlock()
	data:=map[string]interface{}{"meta":"k8s系统关键字","caption":key,"value":key,"score":20}
	AutoCompleteData=append(AutoCompleteData,data)
}
func AddValueAutoComplete(key string){
	mutex.Lock()
	defer mutex.Unlock()
	data:=map[string]interface{}{"meta":"k8s系统关键字-用于值","caption":key,"value":key,"score":20}
	AutoCompleteData=append(AutoCompleteData,data)
}
func DeleteAutoComplete(meta string)  {
	mutex.Lock()
	defer mutex.Unlock()
	for i,v:=range AutoCompleteData{
		if meta==v.(map[string]interface{})["meta"].(string){
			AutoCompleteData= append(AutoCompleteData[:i], AutoCompleteData[i+1:]...)
		}
	}
}
func(this *OtherCtl) autoComplete(c *gin.Context) goft.Json{
	////验证账号
	return gin.H{
		"code":200,
		"data":AutoCompleteData,
	}
}

func(this *OtherCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/autoComplete",this.autoComplete)

}
func(this *OtherCtl)  Name() string{
	 return "OtherCtl"
}