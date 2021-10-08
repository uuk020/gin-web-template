<h1 align="center"> gin-web-template </h1>

<p align="center">gin 二次封装 web 框架. 项目分层设计, 便以开发 web 项目.</p>

### 涉及 go 第三包

| 包名  | 描述 |
|--------|-------|
|gin|web 框架|
|viper|配置管理|
|gorm|orm 框架|
|mysql|数据库|
|redis|缓存数据库|
|jwt|身份认证|
|zap|日志管理器|

### 目录结构设计

```
项目部署目录（或者子目录）
├─cmd                   程序目录
│  └─main.go            应用主入口
│
├─configs               配置目录
│  ├─config.go          应用配置
│
├─initialize            应用初始化
│  ├─config.go          配置初始化
│  ├─database.go        数据库初始化
│  ├─logger.go          日志初始化
│  ├─router.go          路由初始化
│
├─internal              私有应用程序和库代码
│  │─app                应用   
│  │ ├─controller       控制器目录
│  │ ├─middleware       中间件目录
│  │ ├─models           模型目录
│  │ ├─router           路由目录
│  │ ├─service          服务层目录
│  │ ├─static           资源目录
│  │ ├─utils            应用核心目录
│
├─logs                  错误日志
├─go.mod                Go module相关
├─go.sum                Go module相关
├─LINCENSE              授权说明
├─README.md             README 文件
├─settings-dev.yaml     环境变量示例文件
```

### 使用说明
- git clone 此项目, 修改 go.mod 的包名称改成你的
- setting-dev.yaml 重命名为 setting.yaml
- 路由定义统一在 internal/app/router/app.go 编写, 通过 ```utils.AddRoute``` 定义路由, 具体使用可看示例

### 参考
* [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md)
* [10篇带你手摸手封装gin框架](https://juejin.cn/column/6968662583138238478)
* [admin-go](https://github.com/basefas/admin-go)
* [go-clean-template](https://github.com/evrone/go-clean-template)
* [gin 框架文档](https://github.com/gin-gonic/gin#benchmarks)
* [gorm 文档](https://gorm.io/zh_CN/)

### 最后
欢迎提出 issue 和 pull request

### License
MIT