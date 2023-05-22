package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type JwtConfig struct {
	Secret string
	Issuer string
	Expire time.Duration
}

const (
	PORT       = ":8432"
	APP_NAME   = "go-test"
	APP_SECRET = "6YJSuc50uJ18zj45"
	//API过期时间 秒为单位
	API_EXPIRY = "120"
	//log文件路径
	Log_FILE_PATH = "./logs"
	//log文件名称
	LOG_FILE_NAME = "system.log"
	//jwt盐值
	JWTCONFIG_SECRET = "1111"
	//jwt签发人
	JWTCONFIG_ISSUER = "wxb"
	//jwt过期时间
	JWTCONFIG_EXPIRE = time.Hour * 24
)

// InitConfig 初始化配置文件
func InitConfig() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//log.Printf("jwt.secert=%d\n", viper.GetInt("jwt.expires"))
}
