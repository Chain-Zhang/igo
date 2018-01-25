package ilog

import (
	"errors"
	"igo/conf"
)

var AppLog ILog

type ILog interface{
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

func init(){
	level, _ := conf.AppConfig.GetInt("log.level")
	adapter := conf.AppConfig.GetString("log.adapter")
	AppLog, _ = NewILog(adapter, level)
}

func NewILog(adapter string, log_level int) (ILog, error){
    switch adapter{
	case "file":
		return newFileLog(log_level), nil
		default :
		return nil, errors.New("系统为定义该该类型日志")
	}
}