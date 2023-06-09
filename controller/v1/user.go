package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-test/model"
	"go-test/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Person struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// UserLogin 导出的方法（或函数、变量）需要使用大写字母开头的大驼峰写法来表示
func UserLogin(c *gin.Context) {
	//userInfoData, ok := c.Get("userInfo")
	//if !ok {
	//	// 如果获取userInfo失败，则返回错误信息
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	//	return
	//}
	//// 断言
	//userInfo, ok := userInfoData.(middleware.UserInfo)
	//fmt.Println("---->", userInfo.UserId)
	name := c.Query("name")
	password := c.Query("password")
	var result Person
	filter := bson.M{"name": name, "password": utils.MD5(password)}
	//model.UsersCollection.InsertOne(context.Background(), bson.M{"name": name, "password": password})
	model.UsersCollection.FindOne(context.Background(), filter).Decode(&result)
	//fmt.Println("userId==>", result.Id)

	if (Person{} == result) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "用户名或密码错误",
		})
	} else {
		token, _ := utils.CreateToken(result.Id.Hex())
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "登录成功",
			"data": map[string]interface{}{
				"id":       result.Id,
				"name":     result.Name,
				"password": result.Password,
				"token":    token,
			},
		})
	}

}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	var user = &Person{
		Id:        primitive.NewObjectID(),
		Name:      name,
		Password:  utils.MD5(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	model.UsersCollection.InsertOne(context.Background(), user)
	//model.UsersCollection.update({})
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "注册成功",
	})
}
