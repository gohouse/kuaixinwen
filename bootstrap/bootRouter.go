package bootstrap

import (
	"github.com/devfeel/dotweb"
	"net/http"
)

var App *dotweb.DotWeb

func init() {
	//init DotApp
	App = dotweb.New()
	//set log path
	App.SetLogPath("../log")

}

func ServeHTTP() {
	var w http.ResponseWriter
	var r *http.Request
	// Stop here if its Preflighted OPTIONS request
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"*")   //有使用自定义头 需要这个,Action, Module是例子
	}

	if r.Method == "OPTIONS" {
		return
	}

	//w.Write([]byte("hello"))
}