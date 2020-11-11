package aes

import (
	"DesCode/utils"
	"crypto/aes"
	"crypto/cipher"
)

//使用Aes算法对明文进行加密
//三元素 key data mode
func AesEnCrypt(origin []byte ,key []byte)([]byte,error)  {
	block,err:=aes.NewCipher(key)
	if err!=nil {
		return nil,err
	}
	//对明文数据进行尾部填充
    cryptData:=utils.PKCS5EndPadding(origin,block.BlockSize())
    //实例化加密mode
    blockMode :=cipher.NewCBCEncrypter(block,key)
    //加密
    cipherData :=make([]byte,len(cryptData))
    blockMode.CryptBlocks(cipherData,cryptData)
    return cipherData,nil
}