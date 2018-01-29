# kuaixinwen
言简意赅的快餐式新闻网站, gorose orm+dotweb框架快速构建go web网站实战项目
## 在线体验地址

[快新闻](http://kxw.fizzday.net)
[快新闻后台](http://kxw.fizzday.net/#/admin)
## 特色
- 后端采用go编写,基本的增删改查啥都有  
- 前端采用vue+iview+axios+vue-router, 下拉上拉加载更多数据  
- 基本的mvc, 结构化体验  
- 入口统一驱动orm, 路由, 配置
- 开放式后台管理系统(咳咳, 就是不用登录直接可以操作体验, 说明白点就是:懒,没做登录)  

## 开箱体验的正确姿势
- 下载  
```go
go get github.com/gohouse/kuaixinwen
```
- 启动go服务
```go
cd github.com/gohouse/kuaixinwen
go run main.go
```
> 记得修改项目根目录 config 文件夹下的数据库配置

- 安装vue的相关服务
```go
cd github.com/gohouse/kuaixinwen/view
npm install
npm run dev
```

## 有问题是的正确处理姿势
- 1. 提issue

- 2. [点击加入qq群: 470809220](https://jq.qq.com/?_wv=1027&k=5JJOG9E) 解锁新姿势
