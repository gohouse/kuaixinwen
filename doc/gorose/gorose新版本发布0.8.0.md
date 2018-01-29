## gorose orm新版本(0.8.0)发布,新增连接池等特色功能
经过几个日夜, gorose群众多个成员的共同艰苦奋战, 全新版本 0.8.0 发布, 在遵循简单易用的基础之上, 做出了许多改进.  
gorose orm既然号称 gol orm中最风骚的orm, 这次的改进, 也是风骚范儿十足, 下面我们一起来看看改进后的特色吧:  

### 一. 按照开源项目标准重构目录,让更多的人可以自由协作共同开发
```sh
/docs/      ---- 文档目录, 这里包含多个语言的不同使用文档
/drivers/   ---- 不同数据库的驱动目录, 可以自由增加任何其他数据库的目录
/examples/  ---- 使用示例目录, 可以在这里找到大部分的用例
/test/      ---- go testing 自动测试, 包括简单的压力测试
/utils/     ---- 工具包, 放置常用工具函数
/vendor/    ---- 采用glide管理的依赖包
database.go ---- 数据库映射操作的核心文件
glide.yaml  ---- 项目依赖管理的配置文件
gorose.go   ---- 数据库链接,数据库驱动加载核心文件
README.md   ---- 文档说明文件
```
- 调整后的目录, 更加清晰明了.  
- 如果要增加数据库驱动, 只需要在 drivers 目录下添加对应驱动即可  
- 新的文档, 或者翻译文档, 可以直接在 docs 下直接添加  
- 任何人, 都可以直接fork之后, 然后增加内容或修改内容, 发起 pull requests , 合作开发维护清晰无压力    

### 二. 增加了连接池  
采用了官方的连接池解决方案, 使用非常方便, 只需要在配置文件中, 设定对应的连接池参数即可  
在 goroutine 开启的情况下, 数据库使用性能, 有了大大的提升  

### 三. 采用 glide 依赖管理
- glide 依赖管理, 让依赖更加清晰明了   
- 不再爬墙, 缓存了gorose项目依赖的所有包, 再也不用费力找梯子了, 真正的做到了开箱即用   

### 四. 新增了多个用法
比如:  
- 查询一个字段的值, users 表中 id 为 1 的用户的名字叫 fizzday, 只需要 使用 gorose orm 的value方法, 即可以拿到名字 fizzday    
```go
db.Table("users").Where("id",1).Value("name")
```
- 新增了 having 用法:  
```go
db.Table("users").Fields("id, age, sum(money) as sum_money").Group("age").Having("sum_money>10").Get()
```
- 更多请往下看  

-----------
更多 gorose orm 的风骚特性, 请查看项目官方文档 [https://github.com/gohouse/gorose](https://github.com/gohouse/gorose)   

或者 [点击加入qq群: 470809220](https://jq.qq.com/?_wv=1027&k=5JJOG9E) 慢慢撩~~~  

-----------
powered py [fizzday](http://gorose.fizzday.net/)(星期八)
