package configs

// ServerConfig web 服务相关配置结构体
type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MySql       MysqlConfig `mapstructure:"mysql"`
	Redis       RedisConfig `mapstructure:"redis"`
	LogsAddress string      `mapstructure:"logs_address"`
	Timezone    string      `mapstructure:"timezone"`
	Lang        string      `mapstructure:"lang"`
	JWTKey      JWTConfig   `mapstructure:"jwt"`
}

// MysqlConfig mysql 配置结构体
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

// RedisConfig redis 配置结构体
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// JWTConfig jwt 配置
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}