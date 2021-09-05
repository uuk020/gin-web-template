<h1 align="center"> gin-web-template </h1>

<p align="center">gin 二次封装 web 框架. 项目分层设计, 便以开发 web 项目.</p>

### 涉及 go 第三包:

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

| 文件夹  | 描述 |
|--------|-------|
|config|配置|
|controller|控制器|
|dao|操作数据库,给service提供数据|
|forms|验证字段|
|global|全局变量|
|initialize|gin 服务初始化|
|middlewares|中间件|
|models|模型层|
|response|响应格式|
|router|路由|
|service|服务层|
|static|资源文件夹|
|utils|常用功能|
|settings-dev.yaml|配置文件|
|main.go|服务启动|

### 使用说明:
- setting-dev.yaml 重命名为 setting.yaml 或者 修改 initialize/config.go 中的配置文件名

### 参考资料:
* [10篇带你手摸手封装gin框架](https://juejin.cn/column/6968662583138238478)
* [gin 框架文档](https://github.com/gin-gonic/gin#benchmarks)
* [gorm 文档](https://gorm.io/zh_CN/)

### 最后
欢迎提出 issue 和 pull request

### License
MIT