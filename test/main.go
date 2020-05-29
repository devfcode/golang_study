package main

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
	"fmt"
)

func padding(src []byte,blocksize int) []byte {
	n:=len(src)
	padnum:=blocksize-n%blocksize
	pad:=bytes.Repeat([]byte{byte(padnum)},padnum)
	dst:=append(src,pad...)
	return dst
}

func unpadding(src []byte) []byte {
	n:=len(src)
	unpadnum:=int(src[n-1])
	dst:=src[:n-unpadnum]
	return dst
}

func encryptDES(src []byte,key []byte) []byte {
	block,_:=des.NewCipher(key)
	src=padding(src,block.BlockSize())
	blockmode:=cipher.NewCBCEncrypter(block,key)
	blockmode.CryptBlocks(src,src)
	return src
}

func decryptDES(src []byte,key []byte) []byte {
	block,_:=des.NewCipher(key)
	blockmode:=cipher.NewCBCDecrypter(block,key)
	blockmode.CryptBlocks(src,src)
	src=unpadding(src)
	return src
}

func main()  {
	x:=[]byte("长长的头发，黑黑的眼睛。")
	key:=[]byte("12345678")
	x1:=encryptDES(x,key)
	fmt.Println(string(x1))
	x2:=decryptDES(x1,key)
	fmt.Println(string(x2))
}