package util

import (
	"crypto/cipher"
	"bytes"
	"crypto/des"
	"fmt"
	"strings"

	"crypto/md5"
	"encoding/base64"
	"net/url"
)

/*
* md5加密
*/
func Md5(source string, isUpper bool)string{
	buf := []byte(source)
	has := md5.Sum(buf)
	md5Str := fmt.Sprintf("%x", has)
	if isUpper{
		md5Str = strings.ToUpper(md5Str)
	}
	return md5Str
}

/*
* base64编码
*/
func Base64Encode(source string)string{
	buf := []byte(source)
	return base64.StdEncoding.EncodeToString(buf)
}

/*
* base64解码
*/
func Base64Decode(source string) (string, error){
	buf, err := base64.StdEncoding.DecodeString(source)
	if err != nil{
		return "", err
	}
	return string(buf), nil
}

/*
* Url编码
*/
func UrlEncode(source string)string{
	return url.QueryEscape(source)
}

/*
* Url解码
*/
func UrlDecode(source string) (string, error){
	return url.QueryUnescape(source)
}

/*
* DES加密，CBC模式，pkcs5padding，初始向量用key填充
*/
func DesEncrypt(origData, key string)(string, error){
	origBytes := []byte(origData)
	keyBytes := []byte(key)
	block, err := des.NewCipher(keyBytes)
	if err != nil{
		return "", err
	}
	origBytes = pkcs5Padding(origBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, keyBytes)
	crypted := make([]byte, len(origBytes))

	blockMode.CryptBlocks(crypted, origBytes)
	return string(crypted), nil
}

func pkcs5Padding(ciphertext []byte, blockSize int)[]byte{
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func DesDecrypt(crypted, key string)(string, error){
	cryptByts := []byte(crypted)
	keyByts := []byte(key)
	block, err := des.NewCipher(keyByts)
	if err != nil{
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, keyByts)
	origByts := make([]byte, len(cryptByts))
	blockMode.CryptBlocks(origByts, cryptByts)
	origByts = pkcs5UnPadding(origByts)
	return string(origByts), nil
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
	