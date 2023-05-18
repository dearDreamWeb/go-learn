package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 导出的方法（或函数、变量）需要使用大写字母开头的大驼峰写法来表示
func UserLogin(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	fmt.Printf("name: %s,password: %s", name, password)
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "登录成功",
		"data": map[string]interface{}{
			"name":     name,
			"password": password,
		},
	})
}
