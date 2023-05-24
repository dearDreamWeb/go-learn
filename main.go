package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/config"
	"go-test/middleware"
	"go-test/model"
	"go-test/router"
)

func main() {
	config.InitConfig()
	model.InitMongoDB()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(middleware.JwtVerify())
	engine.Use(middleware.LoggerToFile())
	router.InitRouter(engine)
	fmt.Println("开启成功")
	engine.Run(config.ConfigData.PORT)
}
