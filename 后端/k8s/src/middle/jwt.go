package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/common"
)

type jwt struct {

}
func Newjwt() *jwt {
	return &jwt{}
}
func(this *jwt) OnRequest(ctx *gin.Context) error{
	jwt:=ctx.Request.Header.Get("token")
	User, error:=common.ParseToken(jwt)
	goft.Error(error,"用户鉴权失败")
	ctx.Set("user",User)
	return nil
}
func(this *jwt) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}


