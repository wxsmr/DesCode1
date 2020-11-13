package aes

import (
	"DesCode/utils"
	"crypto/aes"
	"crypto/cipher"
)

//使用Aes算法对明文进行加密
//三元素 key data mode
//func AesEnCrypt(origin []byte ,key []byte)([]byte,error)  {
//	block,err:=aes.NewCipher(key)
//	if err!=nil {
//		return nil,err
//	}
//	//对明文数据进行尾部填充
//   cryptData:=utils.PKCS5EndPadding(origin,block.BlockSize())
//   //实例化加密mode
//   blockMode :=cipher.NewCBCEncrypter(block,key)
//   //加密
//   cipherData :=make([]byte,len(cryptData))
//   blockMode.CryptBlocks(cipherData,cryptData)
//   return cipherData,nil
//}
func AesEnCrypt(origin []byte,key []byte)([]byte,error)  {
	block,err:=aes.NewCipher(key)
	if err!=nil {
		return nil,err
	}
	//将明文进行尾部填充
	cryData:=utils.PKCS5EndPadding([]byte(origin),block.BlockSize())
	//实例化一个模型
	mode:=cipher.NewCBCEncrypter(block,key[:block.BlockSize()])
	//加密
	cipherData:=make([]byte,len(cryData))
	mode.CryptBlocks(cipherData,cryData)
	return cipherData,nil
}

func AesDeCrypt(cipherData []byte,key []byte)([]byte,error)  {
	block,err:=aes.NewCipher(key)
	if err!=nil {
		return nil,err
	}
	mode:=cipher.NewCBCDecrypter(block,key[:block.BlockSize()])
	decipherData:=make([]byte,len(cipherData))
	mode.CryptBlocks(decipherData,cipherData)
	decipherData1:=utils.ClearPKCS5EndPadding(decipherData,block.BlockSize())
	return decipherData1,nil
}