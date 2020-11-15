package _DES

import (
	"CryptCode/utils"
	"crypto/cipher"
	"crypto/des"
	_ "go/types"
)

//该函数用于3des算法的加密
func TripleDesEncrypt(orginaltext []byte,key  []byte)([]byte,error)  {
	//三要素:key,data,mode
	//1.实例化一个cipher
	block,err := des.NewTripleDESCipher(key)
	if err  != nil{
		return nil,err
	}
	//2.对明文进行尾部填充
	cryptData := utils.PKCS5EndPadding(orginaltext,block.BlockSize())
   //3.实例化加密模式
   mode :=  cipher.NewCBCEncrypter(block,key[:block.BlockSize()])/*前闭后开区间，实际取了8个*/
	//4，对填充后的明文进行分组加密
	cipherText := make([]byte,len(cryptData))
	mode.CryptBlocks(cipherText,cryptData)
	return cipherText,nil
}
//该函数用于对3des加密后的密文进行解密
 func TripleDesDecrypt  (ciphers []byte,key  []byte)([]byte,error)  {
	//三元素：key、data、mode
	//1.实例化一个cipher
	block,err := des.NewTripleDESCipher(key)
	if err!=nil{
		return nil,err
	}
	//2.不需要对密文进行尾部填充，可直接使用，实例化mode
	 blockMode := cipher.NewCBCDecrypter(block,key[:block.BlockSize()])
     originText := make([]byte,len(ciphers))
	 blockMode.CryptBlocks(originText,ciphers)
     return originText,nil
}

//该函数将对明文进行尾部填充，采用PKCS5方式
//func PKCS5EndPadding(text []byte,size int) []byte {
//	paddingSize := size - len(text)%size
//	paddingText := bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
//	return append(text,paddingText...)
//}
//
//func ZerosEndPadding(text []byte,sizes int)[]byte  {
//	paddingSize := sizes - len(text)%sizes
//	paddingText := bytes.Repeat([]byte{byte(0)},paddingSize)
//    return append(text,paddingText...)
//}
