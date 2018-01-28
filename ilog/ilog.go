package ilog

import (
	"errors"
	"igo/conf"
)

var appLog ILog

type ILog interface{
	debug(v ...interface{})
	debugf(format string, v ...interface{})

	info(v ...interface{})
	infof(format string, v ...interface{})

	warn(v ...interface{})
	warnf(format string, v ...interface{})

	error(v ...interface{})
	errorf(format string, v ...interface{})

	fatal(v ...interface{})
	fatalf(format string, v ...interface{})
}

func init(){
	level, _ := conf.AppConfig.GetInt("log.level")
	adapter := conf.AppConfig.GetString("log.adapter")
	appLog, _ = NewILog(adapter, level)
}

func NewILog(adapter string, log_level int) (ILog, error){
    switch adapter{
	case "file":
		return newFileLog(log_level), nil
		default :
		return nil, errors.New("系统为定义该该类型日志")
	}
}

func Debug(v ...interface{}){
	appLog.debug(v...)
}

func Debugf(format string, v ...interface{}){
	appLog.debugf(format, v...)
}

func Info(v ...interface{}){
	appLog.info(v...)
}

func Infof(format string, v ...interface{}){
	appLog.infof(format, v...)
}

func Warn(v ...interface{}){
	appLog.warn(v...)
}

func Warnf(format string, v ...interface{}){
	appLog.warnf(format, v...)
}

func Error(v ...interface{}){
	appLog.error(v...)
}

func Errorf(format string, v ...interface{}){
	appLog.errorf(format, v... )
}

func Fatal(v ...interface{}){
	appLog.fatal(v...)
}

func Fatalf(format string, v ...interface{}){
	appLog.fatalf(format, v...)
}