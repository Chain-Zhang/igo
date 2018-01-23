package test

import(
	"fmt"

	"igo/util"
)

func Security_test(){
	source := "123456  "
	fmt.Println("md5 encode:", util.Md5(source, false))

	base64Str := util.Base64Encode(source)
	fmt.Println("base 64 encode: ", base64Str)
	base64DecodeStr, err := util.Base64Decode(base64Str)
	if err == nil{
		fmt.Println("base 64 decode: ", base64DecodeStr)
	}
	
	url := "http:www.baidu.com/s?wd=中国"
	urlEncode := util.UrlEncode(url)
	fmt.Println("url encode: ", urlEncode)
	urlDecode, err := util.UrlDecode(urlEncode)
	if err == nil{
		fmt.Println("url decode: ", urlDecode)
	}

	origData := "123456"
	key := "11"
	crypted, err := util.DesEncrypt(origData, key)
	if err != nil{
        fmt.Println("des encrypt error: ", err)
	}else{
		fmt.Println("des encrypt: ", crypted)
	}

	decrypted, err := util.DesDecrypt(crypted, key)
	if err != nil{
        fmt.Println("des decrypt error: ", err)
	}else{
		fmt.Println("des decrypt: ", decrypted)
	}

	crypted, err = util.TripleDesEncrypt(origData, key)
	if err != nil{
        fmt.Println("des encrypt error: ", err)
	}else{
		fmt.Println("triple des encrypt: ", crypted)
	}

	decrypted, err = util.TripleDesDecrypt(crypted, key)
	if err != nil{
        fmt.Println("triple des decrypt error: ", err)
	}else{
		fmt.Println("triple des decrypt: ", decrypted)
	}

	crypted, err = util.AesCBCEncrypte(origData, key)
	if err != nil{
        fmt.Println("aes error: ", err)
	}else{
		fmt.Println("aes encrypt: ", crypted)
	}

	decrypted, err = util.AesCBCDecrypte(crypted, key)
	if err != nil{
        fmt.Println("aes decrypt error: ", err)
	}else{
		fmt.Println("aes decrypt: ", decrypted)
	}
}