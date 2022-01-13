package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/shenyisyn/goft-gin/goft"
	"io/ioutil"
	"math/big"
	rd "math/rand"
	"os"
	"time"
)

//证书操作相关


func(this *Helper) CertData(path string ) []byte{
	f,err:=os.Open(path )
	goft.Error(err)
	defer f.Close()
	b,err:=ioutil.ReadAll(f)
	goft.Error(err)
	return b
}
func(this *Helper) DeleteK8sUser(cn string) error{


	err:=os.Remove(fmt.Sprintf("%s/%s.pem", UserPath,cn))
	if err != nil {
		return err
	}
	err=os.Remove(fmt.Sprintf("%s/%s_key.pem", UserPath,cn))
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


	clientCertFile, err := os.OpenFile(fmt.Sprintf("%s/%s.pem", UserPath,cn), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
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
	clientKeyFile, _ := os.OpenFile(fmt.Sprintf("%s/%s_key.pem", UserPath,cn), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	err = pem.Encode(clientKeyFile, keyPem)
	if err != nil {
		return err
	}
	return  nil
}

