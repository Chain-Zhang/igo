package conf

import (
	"strings"
	"strconv"

	"io"
	"bufio"
	"os"
)

type IniConfig struct{
	ConfigMap  map[string]string
	strcet string
}

func GetIniConfig(filename string)(*IniConfig, error){
	path , err := GetCurrentPath(filename)
	if err != nil{
		return nil, err
	}
	middle := "."
	config :=  new(IniConfig)
	config.ConfigMap = make(map[string]string)
	//打开文件
	file, err := os.Open(path)
	if err != nil{
		return nil, err
	}
	defer file.Close()
	read := bufio.NewReader(file)
	for{
		b, _, err := read.ReadLine()
		if err != nil {
            if err == io.EOF{
				break
			}
			return nil, err
		}
		str := strings.TrimSpace(string(b))
		//配置文件中的注释
		if strings.Index(str, "#") == 0{
			continue
		}
		//配置文件中的前缀处理
		n1 := strings.Index(str, "[")
		n2 := strings.LastIndex(str, "]")
		if n1 > -1 && n2 > -1 && n2 > n1 + 1{
			config.strcet = strings.TrimSpace(str[n1 + 1 : n2])
			continue
		}
		if len(config.strcet) < 1{
			continue
		}
		//
		eqIndex := strings.Index(str, "=")
		if eqIndex < 0{
			continue
		}
		eqLeft := strings.TrimSpace(str[0:eqIndex])
		if len(eqLeft) < 1{
			continue
		}
		eqRight := strings.TrimSpace(str[eqIndex+1:])
		pos := strings.Index(eqRight,"\t#")
		val := eqRight
		if pos > -1{
			val = strings.TrimSpace(eqRight[0:pos])
		}

		pos = strings.Index(eqRight, " #")
		if pos > -1{
			val = strings.TrimSpace(eqRight[0:pos])
		}

		pos = strings.Index(eqRight, "\t//")
		if pos > -1{
			val = strings.TrimSpace(eqRight[0:pos])
		}

		pos = strings.Index(eqRight, " //")
		if pos > -1{
			val = strings.TrimSpace(eqRight[0:pos])
		}
		if len(val) < 1{
			continue
		}
		key := config.strcet + middle + eqLeft
		config.ConfigMap[key] = strings.TrimSpace(val)
	}
	return config, nil
}

func (self *IniConfig) Get(key string)string{
	v, ok := self.ConfigMap[key]
	if ok{
		return v
	}
	return ""
}

func (self *IniConfig) GetString(key string)string{
	return self.Get(key)
}

func (self *IniConfig) GetInt(key string)(int, error){
	return strconv.Atoi(self.Get(key))
}

func (self *IniConfig) GetInt64(key string)(int64, error){
	return strconv.ParseInt(self.Get(key), 10, 64)
}

func (self *IniConfig) GetFloat(key string)(float64, error){
	return strconv.ParseFloat(self.Get(key), 64)
}

func (self *IniConfig) GetBool(key string)(bool, error){
	return strconv.ParseBool(self.Get(key))
}

