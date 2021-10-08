package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/uuk020/gin-web-template/internal/app/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库相关

// InitMysql 初始化 mysql
func InitMysql() {
	mysqlConfig := utils.Settings.MySql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Name, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Mysql 初始化发生错误: " + err.Error())
	}
	utils.MysqlDataBase = db
}

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", utils.Settings.Redis.Host, utils.Settings.Redis.Port)
	utils.RedisDataBase = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: utils.Settings.Redis.Password,
		DB:       0,
	})
	_, err := utils.RedisDataBase.Ping().Result()
	if err != nil {
		panic("Redis 初始化发生错误: " + err.Error())
	}
}
