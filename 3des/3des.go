package _des

import (
	"DesCode/utils"
	"crypto/cipher"
	"crypto/des"
)

//该函数用于实现3des算法的加密
func TripleDesEncrypt(origintext,key[]byte)([]byte,error) {
	//实例化一个cipher
	block,err:=des.NewCipher(key)
	if err!=nil {
		return nil,err
	}
	//对明文进行尾部填充
	//cryptData :=PKCS5EndPadding(origintext,block.BlockSize())
	cryptData :=utils.ZerosEndPadding(origintext,block.BlockSize())
	//实例化加密模式
	mode:=cipher.NewCBCEncrypter(block,key)
	//对填充后的明文进行分组加密
	cipherText := make([]byte,len(cryptData))
	mode.CryptBlocks(cipherText,cryptData)
	return  cipherText,nil
}
//该函数用于对3DES加密的密文进行解密,并返回明文
func TripleDesDecrypt(cipherText []byte,key []byte) ([]byte,error) {
	//三元素:key ,data ,mode
	block,err:=des.NewTripleDESCipher(key)
	if err!=nil {
		return nil,err
	}
	//不需要对密文进行尾部填充,可以直接实例化mode
	blockMode:=cipher.NewCBCDecrypter(block,key)
	originText :=make([]byte,len(cipherText))
	blockMode.CryptBlocks(originText,cipherText)
	return originText,nil
}


//该函数将对明文进行尾部填充,采用PKCS5方式
