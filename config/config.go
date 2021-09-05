package config

// ServerConfig web 服务相关配置结构体
type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MySql       MysqlConfig `mapstructure:"mysql"`
	Redis       RedisConfig `mapstructure:"redis"`
	LogsAddress string      `mapstructure:"logs_address"`
}

// MysqlConfig mysql 配置结构体
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

// RedisConfig redis 配置结构体
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}
