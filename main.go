package main

import (
	_DES "CryptCode/3DES"
	"CryptCode/AES"
	"CryptCode/DES"
	"bytes"
	"fmt"
)

func main() {
	//des:块加密
	//des:key,data,mode
	/*
		key :秘钥
		data:明文，即将加密的明文
		mode：俩种模式，加密与解密
	*/
	//key := []byte("c1906042")
	//data := "夏天的风轻轻吹过"
	//	//加密:crypto
	//	block,err := des.NewCipher(key)
	//	if err!=nil {
	//		panic("初始化错误，请重试")
	//	}
	//	//dst,src
	//	dst := make([]byte,len([]byte(data)))
	//	//加密过程
	//	block.Encrypt(dst,[]byte(data))
	//
	//	fmt.Println("密文为",dst)
	//
	//	//解密
	//	deBlock,err := des.NewCipher(key)
	//	if err != nil {
	//		panic("初始化错误，重试")
	//	}
	//		deData := make([]byte,len(dst))
	//	deBlock.Decrypt(deData,dst)
	//
	//		fmt.Println(string(deData))
	//}
	//对数据加密  DES秘钥长度为8字节、3DES秘钥长度为24
	//	key := []byte("c1906042")
	//	data := "爱若山风过林，竟起无数回响"
	//
	//	//1.得到cipher
	//	block, _ := des.NewCipher(key)
	//	//2.对数据明文进行结尾块填充
	//	paddingData := PCKS5Padding([]byte(data), block.BlockSize())
	//	//3.选择模式
	//	mode := cipher.NewCBCEncrypter(block, key)
	//	//4.加密
	//	dstData := make([]byte, len(paddingData))
	//	mode.CryptBlocks(dstData, paddingData)
	//	fmt.Println("加密后的密文", dstData)
	//
	//	//对数据进行解密
	//	//DES三元素:key,data,mode
	//	key1 := []byte("c1906042")
	//	data1 := dstData//待解密的数据
	//	block1,err := des.NewCipher(key1)
	//	if err!= nil {
	//		panic(err.Error())
	//	}
	//	deMode := cipher.NewCBCDecrypter(block1,key1)
	//	originalData := make([]byte,len(data1))
	//	//分组解密
	//   deMode.CryptBlocks(originalData,data1)
	//	originalData1 :=  utils.ClearPKCS5Padding(originalData,block1.BlockSize())
	//   fmt.Println("解密后的内容",string(originalData1))
	//}
	//DES解密
	fmt.Println("================DES加解密========================")
	key := []byte("20201112")
	data := "仙人抚我顶，结发受长生"
	cipherText, err := DES.DESEnCrypt([]byte(data), key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText, err := DES.DESDeCrypt(cipherText, key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DES解密后：", string(originText))

	//3DES解密
	fmt.Println("============================3DES加解密=================================")
	key1 := []byte("202011122020111220201112") //3des的密钥长度为24字节
	data1 := "似有山河入我怀，却向平川入沧海"
	cipherText1, err := _DES.TripleDesEncrypt([]byte(data1), key1)

	if err != nil {
		fmt.Println("3DES解密失败", err.Error())
		return
	}
	originalText1, err := _DES.TripleDesDecrypt(cipherText1, key1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3des解密后的内容：", string(originalText1))

	//三、AES算法
	fmt.Println("===============================AES加解密======================================")
	key2 := []byte("2020111220201112") //8
	data2 := "南村群童欺我老无力，忍能对面为盗贼"
	cipherText2, err := AES.AESEnCrypt([]byte(data2), key2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("AES加密的数据为：", cipherText2)
	//AES解密
	originalText2, err := AES.AESDecrypt(cipherText2, key2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("AES解密后的内容：", string(originalText2))

	//4：RSA算法的使用
	//data3 := "万事有终，莫问西东"
	//priv, err := rsa.CreatePairKeys()
	//if err != nil {
	//	fmt.Println("生成秘钥对出错：", err.Error())
	//	return
	//}
	////将生成的私钥保存到硬盘的pem文件上
	//rsa.GeneratePemFileByPrivateKey(priv)
	//if err != nil {
	//	fmt.Println("生成私钥pem文件出错", err.Error())
	//}
	//err = rsa.GeneratePemFileByPublicKey(priv.PublicKey)
	//if err != nil {
	//	fmt.Println("生成公钥pem文件出错", err.Error())
	//}

	//对数据进行加密
	//cipherText4, err := rsa.RSAEncrypt(priv.PublicKey, []byte(data3))
	//if err != nil {
	//	fmt.Println("加密算法出错", err.Error())
	//	return
	//}
	//originText3, err := rsa.RSADncrypt(priv, cipherText4)
	//if err != nil {
	//	fmt.Println("加密jiemi出错", err.Error())
	//	return
	//}
	//fmt.Println("RSA结果是", string(originText3))
	////
	//signText4, err := rsa.RSASign(priv, []byte(data3), crypto.MD5)
	//if err != nil {
	//	fmt.Println("RSA签名失败", err.Error())
	//	return
	//}
	//verfyResult, err := rsa.RSAVerify(priv.PublicKey, []byte(data3), signText4)
	//if err != nil {
	//	fmt.Println("签名真正失败", err.Error())
	//}
	//if verfyResult {
	//	fmt.Println("验证成功")
	//} else {
	//	fmt.Println("验证失败")
	//}

//	//六:椭圆曲线数字算法
//	fmt.Println("=========椭圆曲线==========")
//	//①生成秘钥
//	pri,err := ecc.GenerateECDSAKey()
//	if err != nil {
//		fmt.Println("生成ECDSA秘钥失败",err.Error())
//		return
//	}
//	//②准备数据
//	data4 := "清风徐来，水波不兴"
//	//③数字签名
//	r,s,err :=err.ECDSASign(pri,[]byte(data4),r,s)
//
}
//明文数据尾部填充
func PCKS5Padding (text []byte,blockSize int)[]byte{
	//计算要填充的块内容的大小
	paddingSize := blockSize - len(text)%blockSize
    paddingText :=bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
    //fmt.Println("明文尾部追加内容",paddingText)
    return append(text,paddingText...)
}

