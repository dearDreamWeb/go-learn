package config

import "time"

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

//var JWTCONFIG JwtConfig
