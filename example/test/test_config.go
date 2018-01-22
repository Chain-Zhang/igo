package test

import(
	"fmt"
	"igo/conf"
)

func config_test(){
	config,err := conf.NewConfig("ini", "config/app.config")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	strVal := config.GetString("default.string")
	fmt.Println("string value:", strVal)

	intVal, err := config.GetInt("default.int")
	if err != nil{
		fmt.Println("get int value error: ", err.Error())
	}
	fmt.Println("int value: ", intVal)

	int64Val, err := config.GetInt64("default.int64")
	if err != nil{
		fmt.Println("get int64 value error: ", err.Error())
	}
	fmt.Println("int64 value: ", int64Val)
	
	floatVal, err := config.GetFloat("default.float")
	if err != nil{
		fmt.Println("get float value error: ", err.Error())
	}
	fmt.Println("float value: ", floatVal)

	boolVal, err := config.GetBool("default.bool")
	if err != nil{
		fmt.Println("get bool value error: ", err.Error())
	}
	fmt.Println("bool value: ", boolVal)
}