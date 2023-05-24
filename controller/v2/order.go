package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/middleware"
	"go-test/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

// AddOrderParams 不管是传递json还是form传值
// 注意 ，在结构体定义时 首字母必须大写
type Order struct {
	OrderId   string             `bson:"order_id"`
	OrderName string             `bson:"order_name"`
	Price     float64            `bson:"price"`
	UserId    primitive.ObjectID `bson:"userId"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type AddOrderParams struct {
	OrderId   string  `json:"order_id"`
	OrderName string  `json:"order_name"`
	Price     float64 `json:"price"`
}

func AddOrder(c *gin.Context) {
	userInfoData, ok := c.Get("userInfo")
	if !ok {
		// 如果获取userInfo失败，则返回错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// 断言
	userInfo, ok := userInfoData.(middleware.UserInfo)

	var postParams AddOrderParams
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &postParams)
	// String 转换ObjectID
	userId, _ := primitive.ObjectIDFromHex(userInfo.UserId)
	var orderInfo = &Order{
		OrderId:   postParams.OrderId,
		OrderName: postParams.OrderName,
		Price:     float64(postParams.Price),
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	model.OrdersCollection.InsertOne(context.Background(), orderInfo)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "订单" + orderInfo.OrderId + "添加成功",
	})
}

func GetOrder(c *gin.Context) {
	userInfoData, ok := c.Get("userInfo")
	if !ok {
		// 如果获取userInfo失败，则返回错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// 断言
	userInfo := userInfoData.(middleware.UserInfo)

	var results []bson.M
	objectUserId, _ := primitive.ObjectIDFromHex(userInfo.UserId)
	filter := bson.M{"userId": objectUserId}
	cur, err := model.OrdersCollection.Find(context.Background(), filter)
	if err = cur.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println("---->", cur)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "查询成功",
		"data":    results,
	})
}
