package utils

import (
	"crypto/md5"
	"crypto/sha256"
)
//使用md5哈希函数进行hash
func Md5Hash(data []byte)([]byte)  {
	md5Hash := md5.New()
	md5Hash.Write(data)
	return md5Hash.Sum(nil)
}

func Sha256Hash(data []byte)([]byte)  {
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
