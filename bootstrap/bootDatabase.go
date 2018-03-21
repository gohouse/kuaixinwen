package bootstrap

import (
	"github.com/gohouse/gorose"
	"github.com/gohouse/kuaixinwen/config"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var DB gorose.Connection

func init() {
	var err error
	// 加载database
	DB,err = gorose.Open(config.DbConfig)

	if err!=nil{
		fmt.Println(err)
		panic("数据库链接失败")
	}
	errs := DB.Ping()
	if errs!=nil{
		fmt.Println(err)
		panic("数据库链接失败")
	}
}
