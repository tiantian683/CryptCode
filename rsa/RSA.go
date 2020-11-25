package rsa

import (
	"CryptCode/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const RSA_PRIVATE  = "RSA PRIVATE KEY"
const RAS_PUBLIC  = "RSA PUBLIC KEY"
/*
*私钥：
*公钥：
*/

func GenerateKeys(file_name string)error  {
	//1、生成私钥
	pri,err := CreatePairKeys()
	if err != nil {
		return err
	}
	//2、创建私钥文件
	err = GeneratePemFileByPrivateKey(pri,file_name)
	if err!= nil {
		return err
	}
	//3、公钥文件
	err = GeneratePemFileByPublicKey(pri.PublicKey,file_name)
	if err!= nil{
		return err
	}
	return nil
}

//读取pem文件格式的私钥数据
func ReadPemPriKey(file_name string)(*rsa.PrivateKey,error)  {
	blockBytes,err := ioutil.ReadFile(file_name)
	if err != nil {
     return nil,err
	}
	block ,_ :=pem.Decode(blockBytes)

	priBytes :=block.Bytes
	priKey,err := x509.ParsePKCS1PrivateKey(priBytes)
	return priKey,err
}

//读取pem文件格式的公钥数据
func ReadPemPubKey(file_name string)(*rsa.PublicKey,error)  {
	blockBytes,err :=ioutil.ReadFile(file_name)
	if err != nil {
		return nil,err
	}
	block,_:=pem.Decode(blockBytes)
}

func CreatePairKeys() (*rsa.PrivateKey, error) {
	//1.先生成私钥
	var bits int
	flag.IntVar(&bits,"b",2048,"秘钥长度")
	//fmt.Println(bits)
	privateKey,err := rsa.GenerateKey(rand.Reader,bits)
	if err!=nil {
		return nil, err
	}
	//2.根据私钥生成公钥
	//publicKer := privateKey.Public()

	//3.将公钥和私钥返回
   return privateKey,nil
}

//根据给定的私钥数据生成对应的pem文件
func GeneratePemFileByPrivateKey(pri *rsa.PrivateKey,file_name string) (error) {
	//根据PKC1规则，序列化后的私钥
	priStream := x509.MarshalPKCS1PrivateKey(pri)
	//pem文件，此时privatfile文件为空
	privatFile,err:=os.Create("privatekey"+file_name+".pem")//存私钥的生成的文件
	if err!=nil {
		return err
	}
	block := &pem.Block{
		Type:    RSA_PRIVATE,
		Bytes:   priStream,
	}
    //将准备好的格式内容写入到pem文件
	err = pem.Encode(privatFile, block)
	if err!=nil {
		return err
	}
	return nil
}

//根据公钥生成对应的pem文件，持久化存储
func GeneratePemFileByPublicKey(public rsa.PublicKey,file_name string) error {
	//根据PKC1规则，序列化后的私钥
	Stream := x509.MarshalPKCS1PublicKey(&public)
	//pem文件，此时privatfile文件为空
	block := &pem.Block{
		Type:    RAS_PUBLIC,
		Bytes:   Stream,
	}
	pubFile,err:=os.Create("publickey"+file_name+".pem")//存私钥的生成的文件
	if err!=nil {
	return 	pem.Encode(pubFile, block)
	}
	//将准备好的格式内容写入到pem文件
	err = pem.Encode(pubFile, block)
	if err!=nil {
		return err
	}
	return nil
}
//很久用户传入的内容，自动创建公私钥，
func GenerateKeys(file_name string) (*rsa.PrivateKey, error) {
//1.生成私钥
	pri ,err := CreatePairKeys()
	if err != nil {
		return nil,err
	}
	err=GeneratePemFileByPrivateKey(pri,file_name)
		if err != nil {
			return nil,err
		}

	err=GeneratePemFileByPublicKey(pri.PublicKey,file_name)
	return nil, nil
}

//=================第一组，私钥加密，公钥解密========================
//使用rsa算法对数据进行加密
func RSAEncrypt(key rsa.PublicKey,data []byte)([]byte,error)  {
return rsa.EncryptPKCS1v15(rand.Reader,&key,data)
}

func RSADncrypt(private *rsa.PrivateKey,cipher []byte)([]byte,error)  {
	return rsa.DecryptPKCS1v15(rand.Reader,private,cipher)
}

//============================第二种组合：私钥签名，公钥验签=========================
//使用rsa的算法对数据进行数字签名，并返回签名信息
func RSASign(private *rsa.PrivateKey,data []byte,hash crypto.Hash)([]byte,error) {
	//hashed := utils.Md5Hash(data)
	//return rsa.SignPKCS1v15(rand.Reader, private, crypto.MD5, hashed)
	if hash == crypto.MD5 {
		hashed := utils.Md5Hash(data)
		return rsa.SignPKCS1v15(rand.Reader, private, crypto.MD5, hashed)
	} else if hash == crypto.SHA3_256 {
		hashed := utils.Sha256Hash(data)
		return rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA256, hashed)
	} else {
		fmt.Println("不支持此类算法")
	}
	return nil,nil
}
//使用rsa的算法对数据进行签名验证，并返回签名验证的结果
//验证通过返回true
//否则返回false，同时error中有nil
func RSAVerify(pub rsa.PublicKey,data []byte,signText []byte)  (bool,error){
    hashed := utils.Md5Hash(data)
	err := rsa.VerifyPKCS1v15(&pub,crypto.MD5,hashed,signText)
	return err == nil,err
}