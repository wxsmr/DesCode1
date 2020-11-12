package des_two

import (
	"DesCode/utils"
	"crypto/cipher"
	"crypto/des"
)

func DESEnCrypt(data []byte,key []byte)([]byte,error)  {
//三要素:key , data ,mode
	block,err:=des.NewCipher(key)
	if err!=nil {
		return nil,err
	}
	originData:=utils.PKCS5EndPadding(data,block.BlockSize())
	mode:=cipher.NewCBCEncrypter(block,key)
	cipherData:=make([]byte,len(originData))
	mode.CryptBlocks(cipherData,originData)
	return cipherData,nil
}

func DESDeCrypt(data []byte , key []byte)([]byte,error) {
	block,err:=des.NewCipher(key)
	if err!=nil {
		return nil,err
	}
	mode:=cipher.NewCBCDecrypter(block,key)
	originalData:=make([]byte,len(data))
	mode.CryptBlocks(originalData,data)
	originalData = utils.ClearPKCS5EndPadding(originalData,block.BlockSize())
	return originalData,nil
}