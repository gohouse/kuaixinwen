package main

import (
	"github.com/gohouse/utils"
	"github.com/gohouse/kuaixinwen/router"
	. "github.com/gohouse/kuaixinwen/bootstrap"
)

func main()  {
	// 加载数据库
	defer DB.Close()

	// 加载路由
	router.Run(App.HttpServer)

	// 启动web服务
	err := App.StartServer(8888)

	utils.CheckErr(err)
}

func StartServer() error {
	return nil
}
