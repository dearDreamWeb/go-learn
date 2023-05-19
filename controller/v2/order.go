package v2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 不管是传递json还是form传值
// 注意 ，在结构体定义时 首字母必须大写
type AddOrderParams struct {
	OrderId string `json:"orderId"`
}

func AddOrder(c *gin.Context) {
	var postParams AddOrderParams
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &postParams)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "订单" + postParams.OrderId + "添加成功",
	})
}
