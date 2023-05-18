package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/config"
	"go-test/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	router.InitRouter(engine)
	fmt.Println("开启成功")
	engine.Run(config.PORT)
}
