package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt"
	"go-test/config"
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

type MyCustomClaims struct {
	jwt.StandardClaims
	Id string
}

// CreateToken 创建Jwt
func CreateToken(id string) (string, error) {
	claims := MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.JWTCONFIG_EXPIRE).Unix(), // 有效期
			Issuer:    config.JWTCONFIG_ISSUER,                        // 签发人
			IssuedAt:  time.Now().Unix(),                              // 签发时间
		}, id,
	}
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newWithClaims.SignedString([]byte(config.JWTCONFIG_SECRET))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTCONFIG_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//println("===id==>", claims.Id)
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
