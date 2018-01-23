package main

import(
	"fmt"
	"crypto/des"
)

func main(){
	key := "c7546eb6e53a08f098bf0af0aaab3cc2c7546eb6e53a08f098bf0af0aaab3cc2"
	keyByts := []byte(key)
	_, err := des.NewCipher(keyByts)
	if err != nil{
		fmt.Println(err.Error())
	}

	_, err = des.NewCipher(keyByts[0:32])
	if err != nil{
		fmt.Println(err.Error())
	}

	_, err = des.NewCipher(keyByts[0:16])
	if err != nil{
		fmt.Println(err.Error())
	}

	_, err = des.NewCipher(keyByts[0:8])
	if err != nil{
		fmt.Println(err.Error())
	}
}