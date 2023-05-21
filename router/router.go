package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-test/controller/v1"
	v2 "go-test/controller/v2"
)

func InitRouter(r *gin.Engine) {

	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/userLogin", v1.UserLogin)
		GroupV1.Any("/userRegister", v1.UserRegister)
	}

	GroupV2 := r.Group("/v2")
	{
		GroupV2.Any("/addOrder", v2.AddOrder)
		GroupV2.Any("/getOrder", v2.GetOrder)
	}
}
