package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)
//工具函数


//map转map
func(*Helper)ArrayToMap(input []map[string]interface{})map[string]string{
	result:=make(map[string]string)
	for _,data:=range input{
		switch t := data["value"].(type) {
	     // %T prints whatever type t has
		case bool:
			result[data["key"].(string)]=strconv.FormatBool(t)
		case string:
			result[data["key"].(string)]=t
		case float64:
			result[data["key"].(string)]=fmt.Sprintf("%f", t)
		case int:
			result[data["key"].(string)]=strconv.Itoa(t)
		}

	}
	return result
}



//计算md5值
func(this *Helper)  Md5Str(str string ) string   {
	w := md5.New()
	_,_=io.WriteString(w, str)
	return  fmt.Sprintf("%x", w.Sum(nil))
}

//是否相等。就是判断md5
func(this *Helper)  CmIsEq(cm1 map[string]string,cm2 map[string]string) bool{
	return this.Md5Data(cm1)==this.Md5Data(cm2)
}
//把 map 变成md5 string
func(this *Helper)  Md5Data(data map[string]string ) string {
	str:=strings.Builder{}
	for k,v:=range data{
		str.WriteString(k)
		str.WriteString(v)
	}
	return this.Md5Str(str.String())
}


