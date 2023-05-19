package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/model"
	"go-test/utils"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Person struct {
	Name     string
	Password string
}

// 导出的方法（或函数、变量）需要使用大写字母开头的大驼峰写法来表示
func UserLogin(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	var result Person
	filter := bson.M{"name": name, "password": utils.MD5(password)}
	//model.UsersCollection.InsertOne(context.Background(), bson.M{"name": name, "password": password})
	model.UsersCollection.FindOne(context.Background(), filter).Decode(&result)

	fmt.Printf("name: %s,password: %s\n", name, utils.MD5(password))

	if (Person{} == result) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "用户名或密码错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "登录成功",
			"data": map[string]interface{}{
				"name":     result.Name,
				"password": result.Password,
			},
		})
	}

}

func UserRegister(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	model.UsersCollection.InsertOne(context.Background(), bson.M{"name": name, "password": utils.MD5(password)})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "注册成功",
	})
}
