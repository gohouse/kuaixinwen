package router

import (
	"github.com/gohouse/kuaixinwen/controller"
	"github.com/devfeel/dotweb"
)

func Run(Route *dotweb.HttpServer) {
	// 首页测试
	Route.GET("/test", func(ctx dotweb.Context) error{
		return ctx.WriteString("快新闻首页!")
	})
	// 新闻列表
	Route.GET("/newslist", controller.GetNewsList)

	news := Route.Group("news")
	news.GET("getlist",controller.GetNewsList)
	news.GET("getnewsbyid",controller.GetNewsById)
	news.POST("addoreidt",controller.NewsAddOrEdit)
	news.POST("del",controller.NewsDel)
	//Hello World
	//Route.GET("/", ShowIndex)
	//Route.POST("/", TestPost)
	//Route.POST("/user/getuserlist", GetUserList)
}

