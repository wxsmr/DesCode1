package utils

import "bytes"

func PKCS5EndPadding(text[]byte,size int)[]byte {
	//如果刚好是整组,则不用进行尾部添加
	//8字节一组,明文:80字节
	paddingSize :=size-len(text)%size
	paddingText :=bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	return  append(text,paddingText...)

}

func ZerosEndPadding(text[]byte,size int)[]byte {
	paddingSize :=size-len(text)%size
	paddingText :=bytes.Repeat([]byte{byte(0)},paddingSize)
	return  append(text,paddingText...)

}
//该函数用于对解密后的数据进行尾部去除
func ClearPKCS5EndPadding(text []byte,size int)[]byte{
	 //paddingSize:=size-len(text)%size
	 //textSize:=len(text)-paddingSize
	 //获取最后一个元素
	 lastEle:=int(text[len(text)-1])
	 return text[0:len(text)-lastEle]
}