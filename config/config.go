package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type ConfigData struct {
	PORT string
	//log文件路径
	LOG_FILE_PATH string
	//log文件名称
	LOG_FILE_NAME string
	//jwt盐值
	JWTCONFIG_SECRET string
	//jwt签发人
	JWTCONFIG_ISSUER string
	//jwt过期时间
	JWTCONFIG_EXPIRE time.Duration
}

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

func Config() ConfigData {
	return ConfigData{
		PORT:             viper.GetString("app.port"),
		LOG_FILE_PATH:    viper.GetString("log.filePath"),
		LOG_FILE_NAME:    viper.GetString("log.fileName"),
		JWTCONFIG_SECRET: viper.GetString("jwt.secert"),
		JWTCONFIG_ISSUER: viper.GetString("jwt.issuer"),
		JWTCONFIG_EXPIRE: time.Duration(viper.GetInt("jwt.expires")),
	}
}
