package conf

import (
	"strings"
	"errors"

	"path/filepath"
)

type Config interface{
	GetString(string) string
	GetInt(string)(int, error)
	GetInt64(string)(int64, error)
	GetFloat(string)(float64, error)
	GetBool(string)(bool, error)
}

func NewConfig(adapter, filename string) (Config, error){
	path, err := GetCurrentPath(filename)
	if err != nil{
		return nil, err
	}
	switch adapter{
	case "ini":
		return GetIniConfig(path)
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}

func GetCurrentPath(filename string) (path string, err error ){
    path, err = filepath.Abs(filename)
    if err != nil {
        return
    }
    path = strings.Replace(path, "\\", "/", -1)
    path = strings.Replace(path, "\\\\", "/", -1)
    return 
}