package main

import (
	"github.com/gohouse/kuaixinwen/router"
	. "github.com/gohouse/kuaixinwen/bootstrap"
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("start...,port:8888, visit: http://localhost:8888")

	// 解决跨域问题
	//ServeHTTP()
	//var w http.ResponseWriter
	//var r *http.Request
	////LDNS(w, req)
	//lvtao(w, r)

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

func StartServer() error {
	return nil
}
func LDNS(w http.ResponseWriter, req *http.Request) {

	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if req.Method == "OPTIONS" {
		return
	}
	//// 响应http code
	//w.WriteHeader(200)
	//query := strings.Split(req.Host, ".")
	//value, err := ldns.RAMDBMgr.Get(query[0])
	//fmt.Println("Access-Control-Allow-Origin", "*")
	//if err != nil {
	//	io.WriteString(w, `{"message": ""}`)
	//	return
	//}
	//
	//io.WriteString(w, value)
}

func lvtao(w http.ResponseWriter, r *http.Request) {
	//这儿可以加上一些判断,比如根据域名,而不是*
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		return
	}
	fmt.Fprintln(w, "hi,taotao")
}
