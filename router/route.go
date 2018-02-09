package router

import (
	// 引入控制器
	"github.com/gohouse/kuaixinwen/controller"
	// 引入dotweb
	"github.com/devfeel/dotweb"
	// 引入dotweb中间件
	"github.com/devfeel/middleware/cors"
)

func Run(Route *dotweb.HttpServer) {
	// 设置cors选项中间件, 并使用默认的跨域配置
	option := cors.NewConfig().UseDefault()

	// 首页测试 /
	Route.GET("/", func(ctx dotweb.Context) error {
		return ctx.WriteString("快新闻首页!")
	})
	// json返回测试 /json
	Route.GET("/json", func(ctx dotweb.Context) error {
		return ctx.WriteJson("浪里个浪...")
	}).Use(cors.Middleware(option))

	// 前台展示列表 /getnewslist
	Route.GET("/getnewslist", controller.GetNewsList).Use(cors.Middleware(option))
	// 获取单条新闻信息 /getnewsbyid
	Route.GET("/getnewsbyid", controller.GetNewsById)

	// 后台管理
	admin := Route.Group("/admin").Use(cors.Middleware(option))
	// 新闻列表 /admin/newsdel
	admin.GET("/getnewslist", controller.GetNewsList)
	// 获取单条新闻信息 /admin/getnewsbyid
	admin.GET("/getnewsbyid", controller.GetNewsById)
	// 新闻增加或修改, 根据是否传了id来判定, 传了id就修改, 否则增加 /admin/newsaddoredit
	admin.POST("/newsaddoredit", controller.NewsAddOrEdit)
	// 新闻删除 /admin/newsdel
	admin.POST("/newsdel", controller.NewsDel)
}
