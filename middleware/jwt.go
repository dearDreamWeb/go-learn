package middleware

import (
	"github.com/gin-gonic/gin"
	"go-test/utils"
	"net/http"
	"strings"
)

// UserInfo 中间件传递的值
type UserInfo struct {
	UserId string
}

var whiteList = []string{"/v1/userLogin", "/v1/userRegister"}

// JwtVerify 校验token，
// 从header 中Authorization获取token，格式是 Bearer + 空格 + token值 /*
func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUrl := c.Request.URL.Path
		if utils.ArrayIncludes(whiteList, requestUrl) {
			c.Next()
			return
		}
		authorizationString := c.Request.Header.Get("Authorization")
		arr := strings.Fields(authorizationString)
		if len(arr) <= 1 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"msg":     "token失效",
			})
			c.Abort()
			return
		}
		res, err := utils.ParseToken(arr[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"msg":     "token失效",
			})
			c.Abort()
			return
		}
		data := UserInfo{
			UserId: res.Id,
		}
		c.Set("userInfo", data)
		println("res===>", res)
		c.Next()
	}
}
