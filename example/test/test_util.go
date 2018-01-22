package test

import(
	"fmt"

	"igo/util"
)

func Security_test(){
	source := "123456"
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
	key := "hdhhdhdmZtmcOlmT"
	crypted, err := util.DesEncrypt(origData, key)
	if err != nil{
        fmt.Println("des encrypt error: ", err.Error())
	}else{
		fmt.Println("des encrypt: ", crypted)
	}
}