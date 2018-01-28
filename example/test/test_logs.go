package test

import(
	"igo/ilog"
	"fmt"
)

func TestLog(){
	ilog.Debug("app log debug message")
	ilog.Info("app log info message")
	ilog.Warn("app log warn message")
	ilog.Error("app log error message")
	ilog.Fatal("app log fatal message")

	var str string
	fmt.Scan(&str)
}