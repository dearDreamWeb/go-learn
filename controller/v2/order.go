package v2

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 不管是传递json还是form传值
// 注意 ，在结构体定义时 首字母必须大写
type AddOrderParams struct {
	Order string `json:"order"`
}

func AddOrder(c *gin.Context) {
	var postParams map[string]string
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &postParams)
	fmt.Println("-----", c.PostForm("orderId"))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "订单" + postParams["orderId"] + "添加成功",
	})
}
