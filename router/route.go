package router

import (
	"github.com/gohouse/kuaixinwen/controller"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/cors"
	"github.com/gohouse/utils"
)

func Run(Route *dotweb.HttpServer) {
	// 设置cors选项中间件
	option := cors.NewConfig().UseDefault()

	// 首页测试
	Route.GET("/", func(ctx dotweb.Context) error{
		return ctx.WriteString("快新闻首页!")
	})
	// json返回测试
	Route.GET("/json", func(ctx dotweb.Context) error{
		return ctx.WriteJson(utils.SuccessReturn())
	}).Use(cors.Middleware(option))

	// 前台展示列表
	Route.GET("/getnewslist",controller.GetNewsList).Use(cors.Middleware(option))

	// 后台管理
	admin := Route.Group("/admin").Use(cors.Middleware(option))
	admin.GET("/getnewslist",controller.GetNewsList)
	admin.GET("/getnewsbyid",controller.GetNewsById)
	admin.POST("/newsaddoreidt",controller.NewsAddOrEdit)
	admin.POST("/newsdel",controller.NewsDel)
}

