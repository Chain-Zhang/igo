package util

import (
	"fmt"
	"strings"
	"crypto/md5"
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