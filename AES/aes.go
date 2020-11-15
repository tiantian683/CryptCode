package AES

import (
	"CryptCode/utils"
	"crypto/aes"
	"crypto/cipher"
)

//使用AES算法对明文进行加密
func AESEnCrypt(origin []byte, key []byte) ([]byte, error) {
	//三元素：key、data、mode
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//对明文数据进行尾部填充
	cryptData := utils.PKCS5EndPadding(origin, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	cipherData := make([]byte, len(cryptData))
	blockMode.CryptBlocks(cipherData, cryptData)
	return cipherData, nil
}

//AES解密
func AESDecrypt(data,key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	originText := make([]byte, len(data))
	mode.CryptBlocks(originText,data)
	originText = utils.ClearPKCS5Padding(originText, block.BlockSize())
	return originText,nil
}
