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
