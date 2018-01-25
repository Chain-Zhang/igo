package util

import (
	"os"
	"fmt"

	"bytes"
	"strings"
	"errors"

	"crypto/des"
	"crypto/aes"
	"crypto/cipher"

	"encoding/base64"
	"net/url"
)

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
	keyBytes := getDESKey(key)
	block, err := des.NewCipher(keyBytes)
	if err != nil{
		return "", err
	}
	origBytes = pkcs5Padding(origBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, keyBytes)
	crypted := make([]byte, len(origBytes))

	blockMode.CryptBlocks(crypted, origBytes)
	return Base64Encode(string(crypted)), nil
}

/*
* DES解密，CBC模式，pkcs5padding，初始向量用key填充
*/
func DesDecrypt(crypted, key string)(string, error){
	crypted, _ = Base64Decode(crypted)
	cryptByts := []byte(crypted)
	keyByts := getDESKey(key)
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

/*
* 三重DES加密，CBC模式，pkcs5padding，初始向量用key填充
*/
func TripleDesEncrypt(origData, key string)(string, error){
	origBytes := []byte(origData)
	keyBytes := getTripleDESKey(key)
	block, err := des.NewTripleDESCipher(keyBytes)
	if err != nil{
		return "", err
	}
	origBytes = pkcs5Padding(origBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:8])
	crypted := make([]byte, len(origBytes))

	blockMode.CryptBlocks(crypted, origBytes)
	return Base64Encode(string(crypted)), nil
}

/*
* 三重DES解密，CBC模式，pkcs5padding，初始向量用key填充
*/
func TripleDesDecrypt(crypted, key string)(string, error){
	crypted, _ = Base64Decode(crypted)
	cryptByts := []byte(crypted)
	keyByts := getTripleDESKey(key)
	block, err := des.NewTripleDESCipher(keyByts)
	if err != nil{
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, keyByts[:8])
	origByts := make([]byte, len(cryptByts))
	blockMode.CryptBlocks(origByts, cryptByts)
	origByts = pkcs5UnPadding(origByts)
	return string(origByts), nil
}

/*
* AES加密，CBC模式，pkcs5padding，初始向量用key填充
*/
func AesCBCEncrypte(origData, key string) (string, error) {
	origByts := []byte(origData)
	keybytes := getAESKey(key)
	plaintext := pkcs5Padding(origByts, aes.BlockSize)
	block, err := aes.NewCipher(keybytes[:aes.BlockSize])
	if err != nil{
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, keybytes[:aes.BlockSize])
	crypted := make([]byte, len(plaintext))
	mode.CryptBlocks(crypted, plaintext)
	return Base64Encode(string(crypted)), nil
}

/*
* AES解密，CBC模式，pkcs5padding，初始向量用key填充
*/
func AesCBCDecrypte(crypted string, key string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "error string:%s key:%s err:%v\n", crypted, key, err)
		}
	}()

	keybytes := getAESKey(key)
	crypted, err := Base64Decode(crypted)
	if err != nil {
		return "", errors.New("crypted data format error")
	}
    cryptedData := []byte(crypted)
	block, err := aes.NewCipher(keybytes[:aes.BlockSize])
	if err != nil{
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, keybytes[:aes.BlockSize])

	decryptedData := make([]byte, len(cryptedData))
	mode.CryptBlocks(decryptedData, cryptedData)
	cryptedData = pkcs5UnPadding(decryptedData)
	return strings.TrimSpace(string(decryptedData)), nil
}

func pkcs5Padding(ciphertext []byte, blockSize int)[]byte{
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func getDESKey(key string)[]byte{
	key = Md5(key, false)
	keyBytes := []byte(key)
	return keyBytes[0:8]
}

func getTripleDESKey(key string)[]byte{
	key = Md5(key, false)
	keyBytes := []byte(key)
	return keyBytes[0:24]
}

func getAESKey(key string) []byte {
    key = Md5(key, false)
    keyLen := len(key)
    arrKey := []byte(key)
    if keyLen >= 32 {
        return arrKey[:32]
    }
    if keyLen >= 24 {
        return arrKey[:24]
    }
    return arrKey[:16]
}