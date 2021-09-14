package utils

import (
	"github.com/go-redis/redis"
	"github.com/uuk020/gin-web-template/configs"
	"gorm.io/gorm"
)
// 全局变量
var (
	// Settings 系统配置
	Settings configs.ServerConfig
	// RouterGroup 路由规则
	RouterGroup map[string][]*Route
	// MysqlDataBase 全局变量
	MysqlDataBase *gorm.DB
	// RedisDataBase 全局变量
	RedisDataBase *redis.Client
)

const (
	// StaticResources 静态资源路径, 相对 main.go 文件所在的
	StaticResources = "./app/static/"
)
