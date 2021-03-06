package main

import (
	"DesCode/Des"
	"DesCode/aes"
	"DesCode/des_two"
	"fmt"
)

//
//import (
//	"bytes"
//	"crypto/cipher"
//	"crypto/des"
//	"fmt"
//)

func main() {

	//key:密钥
	//data:明文,即加密的明文
	//mode:两种模式,加密和解密
	//DES密钥为8字节,3Des为24字节
	//key := []byte("c1906041")

	//data := "遇贵人先立业"

	//加密:crypto
	//block,err := des.NewCipher(key)
	//if err!=nil {
	//		panic("初始化密码错误,请重试!")
	//	}
	//	dst:=make([]byte,len([]byte(data)))
	//	//加密过程
	//	block.Encrypt(dst,[]byte(data))
	//	fmt.Println("加密后的内容",dst)
	//
	//	//解密
	// deBlock,err:=des.NewCipher(key)
	//  if err!=nil {
	//	panic("初始化密码错误,请重试!")
	//	}
	//	deData:=make([]byte,len(dst))
	//
	//  deBlock.Decrypt(deData,dst)
	//	//对数据明文进行结尾块填充
	//	paddingData:=Endpadding([]byte(data),block.BlockSize())
	//	//选择模式
	//	mode:=cipher.NewCBCEncrypter(block,key)
	//	//4.加密
	//	dstData:=make([]byte,len(paddingData))
	//	mode.CryptBlocks(dstData,paddingData)
	//	fmt.Println(dstData)
	//
	//
	//	//对数据进行解密
	//	//DES三元素:key,data,mode
	//	key1:=[]byte("C1906041")
	//	data1:=dstData//待解密的数据
	//	block1,err:=des.NewCipher(key1)
	//	if err!=nil {
	//		panic(err.Error())
	//	}
	//	deMode:=cipher.NewCBCDecrypter(block1,key1)
	//	originalData:=make([]byte,len(data1))
	//	//分组解密
	//	deMode.CryptBlocks(originalData,data1)
	//	originaData:=utils.ClearPKCS5EndPadding(originalData,block1.BlockSize())
	//	fmt.Println("解密后的内容",string(originaData))
	//}
	////明文数据尾部填充
	//func Endpadding(text []byte,blockSize int)[]byte  {
	//	//计算要填充块的大小
	//	paddingSize :=blockSize -len(text)%blockSize
	//	paddingText :=bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	//	return append(text,paddingText...)
	//}
	//func main() {
	//	key:=[]byte("abcdefghijklmnopqrstuvwx")
	//	str:="我爱GO语言"
	//	result,_:=Des.TripleDesEncrypt([]byte(str),key)
	//	fmt.Printf("加密后的数据:%x\n",result)
	//}
	key:=[]byte("20201112")
	data:="龚江华憨批hghuhhj"
	cipherText,err:=des_two.DESEnCrypt([]byte(data),key)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	originText,err:=des_two.DESDeCrypt(cipherText,key)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("解密:后的数据",string(originText))

	//3DES加密
	//3DES算法密钥的长度:24字节,固定不变
	//DES和3DES算法
	key1:=[]byte("202011122020111220201112")//3des的密钥长度是24字节
	fmt.Println(key1)
	key2:=make([]byte,16)
	key2 =append(key2,[]byte("20201112")...)

	data1:="穷在闹市无人问"
	cipherText1,err:=Des.TripleDesEncrypt([]byte(data1),key2)
	if err!=nil {
		fmt.Println("3DES加密失败",err.Error())
		return
	}
	originalText1,err:=Des.TripleDesDecrypt(cipherText1,key2)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("解密后:",string(originalText1))


	//AES算法
	key3:=[]byte("2020111220201112")
	data2:="江华我是你爹"

	cipherText2,err:=aes.AesEnCrypt([]byte(data2),key3)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	origin,err:=aes.AesDeCrypt(cipherText2,key3)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("解密后的数据:",string(origin))
}