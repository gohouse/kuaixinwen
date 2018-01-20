package bootstrap

import (
	"github.com/gohouse/gorose"
	"github.com/gohouse/kuaixinwen/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB gorose.Database

func init() {
	// 加载database
	DB = gorose.Open(config.DbConfig, "mysql_dev")
	DB.Ping()
}
