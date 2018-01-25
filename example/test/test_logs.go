package test

import(
	"igo/ilog"
	"fmt"
)

func TestLog(){
	ilog.AppLog.Debug("app log debug message")
	ilog.AppLog.Info("app log info message")
	ilog.AppLog.Warn("app log warn message")
	ilog.AppLog.Error("app log error message")
	ilog.AppLog.Fatal("app log fatal message")

	var str string
	fmt.Scan(&str)
}