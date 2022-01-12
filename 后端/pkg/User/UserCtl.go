package User

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type UserCtl struct{

}
func NewUserCtl() *UserCtl {
  return &UserCtl{}
}
func(this *UserCtl) sendEmail(c *gin.Context) goft.Json{

	r:=EmailCheck{}
	goft.Error(c.ShouldBindJSON(&r),"邮箱不正确")
	//发送邮件
	if r.Email=="test@ng2-oa.com"{
		goft.Throw("游客邮箱，请使用验证码123456登陆",400,c)
	}
	//发送邮箱
	//查询邮箱是否存在

	return gin.H{
		"code":200,
	}
}

func(this *UserCtl) doLogin(c *gin.Context) goft.Json{
	r:=LoginUser{}
	goft.Error(c.ShouldBindJSON(&r),"参数错误")
	//发送邮件
	fmt.Println(r)
	if r.Email=="test@ng2-oa.com"&&r.Code=="123456"{
		token, e:=GenerateToken(r.Email,"Gust")
		goft.Error(e,"系统异常")
		return gin.H{
			"code":200,
			"token":token,
		}
	}
	////验证账号
	return gin.H{
		"code":400,
	}
}

func(this *UserCtl)  Build(goft *goft.Goft){
	goft.Handle("POST","/email",this.sendEmail)
	goft.Handle("POST","/login",this.doLogin)
	//goft.HandleWithFairing("GET","/test",this.test,middle.Newjwt())
}
func(this *UserCtl)  Name() string{
	 return "UserCtl"
}