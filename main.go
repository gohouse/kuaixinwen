package main

import (
	"github.com/gohouse/kuaixinwen/router"
	. "github.com/gohouse/kuaixinwen/bootstrap"
	"fmt"
)

func main()  {
	fmt.Println("start...,port:8888, visit: http://localhost:8888")

	// 加载数据库
	//defer DB.Close()

	// 加载路由
	router.Run(App.HttpServer)

	// 启动web服务
	err := App.StartServer(8888)

	if err!=nil{
		fmt.Println(err)
	}
}
