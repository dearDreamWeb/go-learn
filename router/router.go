package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-test/controller/v1"
	"go-test/utils"
	"net/url"
	"strconv"
)

func InitRouter(r *gin.Engine) {
	r.GET("/sn", SignDem)

	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/userLogin", v1.UserLogin)
	}
}

func SignDem(c *gin.Context) {
	ts := strconv.FormatInt(utils.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":     []string{"test"},
		"password": []string{"123"},
		"ts":       []string{ts},
	}
	res["sn"] = utils.CreateSign(params)
	res["ts"] = ts
	utils.RetJson("200", "", res, c)
}
