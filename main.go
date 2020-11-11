package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main()  {

	//key:密钥
	//data:明文,即加密的明文
	//mode:两种模式,加密和解密

	key :=[]byte("c1906041")

	data:="遇贵人先立业"

	//加密:crypto
	//cipher
	//block,err := des.NewCipher(key)
	//if err!=nil {
	//	panic("初始化密码错误,请重试!")
	//}
	//dst:=make([]byte,len([]byte(data)))
	//加密过程
	//block.Encrypt(dst,[]byte(data))
	//fmt.Println("加密后的内容",dst)

	//解密
	//deBlock,err:=des.NewCipher(key)
	//if err!=nil {
	//	panic("初始化密码错误,请重试!")
	//}
	//deData:=make([]byte,len(dst))
	//deBlock.Decrypt(deData,dst)
	//fmt.Println(string(deData))



	block,_ := des.NewCipher(key)
	//对数据明文进行结尾块填充
	paddingData:=Endpadding([]byte(data),block.BlockSize())
	//选择模式
	mode:=cipher.NewCBCEncrypter(block,key)
	//4.加密
	dstData:=make([]byte,len(paddingData))
	mode.CryptBlocks(dstData,paddingData)
	fmt.Println(dstData)

	//对数据进行解密
	//DES三元素:key,data,mode
	key1:=[]byte("C1906041")
	data1:=dstData//待解密的数据
	block1,err:=des.NewCipher(key1)
	if err!=nil {
		panic(err.Error())
	}
	deMode:=cipher.NewCBCDecrypter(block1,key1)
	originalData:=make([]byte,len(data1))
	//分组解密
	deMode.CryptBlocks(originalData,data1)
	fmt.Println("解密后的内容",string(originalData))
}
//明文数据尾部填充
func Endpadding(text []byte,blockSize int)[]byte  {
	//计算要填充块的大小
	paddingSize :=blockSize -len(text)%blockSize
	paddingText :=bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	return append(text,paddingText...)
}