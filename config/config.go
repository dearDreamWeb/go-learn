package config

import "time"

type JwtConfig struct {
	Secret string
	Issuer string
	Expire time.Duration
}

const (
	PORT          = ":8432"
	APP_NAME      = "go-test"
	APP_SECRET    = "6YJSuc50uJ18zj45"
	API_EXPIRY    = "120"
	Log_FILE_PATH = "./logs"
	LOG_FILE_NAME = "system.log"
	//JWTCONFIG     = JwtConfig{
	//	Secret: "",
	//	Issuer: "",
	//	Expire: 1000 * 60,
	//}
)

//var JWTCONFIG JwtConfig
