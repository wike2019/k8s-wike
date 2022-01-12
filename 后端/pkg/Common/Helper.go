package Common

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/shenyisyn/goft-gin/goft"
	"io"
	"io/ioutil"
	"math/big"
	rd "math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)
const UserPath ="./user"
const CaKeyPath = "./ca.key"
const CaCrtPath = "./ca.crt"
type ItemsPage struct {
	Total        int //一共多少条
	Current         int //当前页
	Size     int // 页尺寸
	PageNum int //一共多少页
	Data []interface{}  //数据
	Ext interface{} //扩展信息，方便插入值 给前端用
}
func(this *ItemsPage) SetExt(ext interface{}) *ItemsPage{
	this.Ext=ext
	return this
}
//@Component
type Helper struct {
//	PodMap *Pod.PodMapStruct `inject:"-"`
}
func NewHelper() *Helper {
	return &Helper{}
}
//字符串转int
func(*Helper) StrToInt(str string,def int) int {
	ret,err:=strconv.Atoi(str)
	if err!=nil{
		return def
	}
	return ret
}
//分页 资源
func(*Helper)   PageResource(current ,size  int,list []interface{}) *ItemsPage{
	total:=len(list)
	if size==0 || size>total{
		size=5 //默认 每页5个
	}
	if current<=0{
		current=1
	}
	pageInfo:=&ItemsPage{Total:total,Size:size}
	//计算总页数
	pageNum:=1
	if pageInfo.Total>size{
		pageNum=pageInfo.Total/size
		if pageInfo.Total %size!=0{
			pageNum++
		}
	}
	if current>pageNum{
		current=1
	}
	pageInfo.Current=current //重新赋值Current ----当前页
	newSet:=make([]interface{},0) //构建一个新的 切片

	if current*size>pageInfo.Total{
		newSet=append(newSet,list[(current-1)*size:]...)
	}else {
		//fmt.Println((current-1)*size,":",size)
		//  1 ,2,3,4,5,6
		// [) 左闭右开
		// 0,2   list[0:2]
		newSet=append(newSet,list[(current-1)*size:(current-1)*size+size]...)
	}
	//重新整理赋值
	pageInfo.Data=newSet
	pageInfo.PageNum=pageNum
	return pageInfo
}

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


func(this *Helper) CertData(path string ) []byte{
	f,err:=os.Open(path )
	goft.Error(err)
	defer f.Close()
	b,err:=ioutil.ReadAll(f)
	goft.Error(err)
	return b
}
func(this *Helper) DeleteK8sUser(cn string) error{


	err:=os.Remove(fmt.Sprintf("%s/%s.pem",UserPath,cn))
	if err != nil {
		return err
	}
	err=os.Remove(fmt.Sprintf("%s/%s_key.pem",UserPath,cn))
	if err != nil {
		return err
	}
	return  nil
}
func(this *Helper) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func(this *Helper) GenK8sUser(cn string,o string) error{
	ok,err:=this.PathExists(UserPath)
	fmt.Println(ok)
	if !ok{
		os.MkdirAll(UserPath,0755)
	}
	caFile, err := ioutil.ReadFile(CaCrtPath)
	if err != nil {
		return err
	}
	caBlock, _ := pem.Decode(caFile)

	caCert, err := x509.ParseCertificate(caBlock.Bytes) //ca 证书对象
	if err != nil {
		return err
	}
	//解析私钥
	keyFile, err := ioutil.ReadFile(CaKeyPath)
	if err != nil {
		return err
	}
	keyBlock, _ := pem.Decode(keyFile)
	caPriKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes) //是要对象
	if err != nil {
		return err
	}
	//---------- ---------------------  ---------------------------------
	//构建证书模板
	certTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(rd.Int63()), //证书序列号
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{o},
			//OrganizationalUnit: []string{"可填课不填"},
			Province:           []string{"beijing"},
			CommonName:         cn,//CN
			Locality:           []string{"beijing"},
		},
		NotBefore:             time.Now(),//证书有效期开始时间
		NotAfter:              time.Now().AddDate(1, 0, 0),//证书有效期
		BasicConstraintsValid: true, //基本的有效性约束
		IsCA:                  false,   //是否是根证书
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, //证书用途(客户端认证，数据加密)
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		EmailAddresses:        []string{"200569525@qq.com"},
	}

	//生成公私钥--秘钥对
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	//创建证书 对象
	clientCert, err := x509.CreateCertificate(rand.Reader, certTemplate, caCert, &priKey.PublicKey, caPriKey)
	if err != nil {
		return err
	}

	//编码证书文件和私钥文件
	clientCertPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCert,
	}


	clientCertFile, err := os.OpenFile(fmt.Sprintf("%s/%s.pem",UserPath,cn), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	err = pem.Encode(clientCertFile, clientCertPem)
	if err != nil {
		return err
	}

	buf := x509.MarshalPKCS1PrivateKey(priKey)
	keyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: buf,
	}
	clientKeyFile, _ := os.OpenFile(fmt.Sprintf("%s/%s_key.pem",UserPath,cn), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	err = pem.Encode(clientKeyFile, keyPem)
	if err != nil {
		return err
	}
	return  nil
}
