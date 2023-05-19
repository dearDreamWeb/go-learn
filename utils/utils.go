package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/config"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

func GetTimeUnix() int64 {
	return time.Now().Unix()
}

func MD5(str string) string {
	b := []byte(str)
	s := []byte("gt4si3tbrl8udpla4dlv9wta")
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	str := ""
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	sign := MD5(MD5(str) + MD5(config.APP_NAME+config.APP_SECRET))
	return sign
}

func VerifySign(c *gin.Context) {
	var method = c.Request.Method
	var ts int64
	var sn string
	var req url.Values

	if method == "GET" {
		req = c.Request.URL.Query()
		sn = c.Query("sn")
		ts, _ = strconv.ParseInt(c.Query("time"), 10, 64)
	} else if method == "POST" {
		c.Request.ParseForm()
		req = c.Request.PostForm
		sn = c.PostForm("sn")
		ts, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
	} else {
		RetJson("500", "Illegal requets", "", c)
		return
	}
	exp, _ := strconv.ParseInt(config.API_EXPIRY, 10, 64)

	//	验证过期时间
	if ts > GetTimeUnix() || GetTimeUnix()-ts >= exp {
		RetJson("500", "Ts Error", "", c)
		return
	}
	if sn == "" || sn != CreateSign(req) {
		RetJson("500", "Sn Error", "", c)
		return
	}
}

func RetJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}
