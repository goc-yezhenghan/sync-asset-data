package main

import (
	"fmt"
	"syncData/initServe"
	"syncData/operation"
)

func main() {

	// 初始化viper
	initServe.InitConfig()

	// 初始化数据库
	initServe.InitMySQL()

	// 操作数据
	operation.TransferData()

	fmt.Println("Exec successfully")
}
