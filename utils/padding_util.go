package utils

import "bytes"

//该函数将对明文进行尾部填充，采用PKCS5方式
func PKCS5EndPadding(text []byte,size int) []byte {
	paddingSize := size - len(text)%size
	paddingText := bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	return append(text,paddingText...)
}

//该函数用于对解密后的数据进行尾部去除
func ClearPKCS5Padding(text []byte,size int) []byte  {
	lastEle := int(text[len(text)-1])
	return text[0:len(text)-lastEle]
}

//该函数用于对明文进行尾部填充，采用Zeros方式
func ZerosEndPadding(text []byte,size int)[]byte  {
	if len(text)%size == 0{
		return text
	}
	paddingSize := size - len(text)%size
	paddingText := bytes.Repeat([]byte{byte(0)},paddingSize)
	return append(text,paddingText...)
}
