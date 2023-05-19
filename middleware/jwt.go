package middleware

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type MyCustomClaims struct {
	jwt.StandardClaims
	Uid string
}

// 创建Jwt
func CreateToken(uid string) (string, error) {

	claims := MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: 1000 * 30, // 有效期
			Issuer:    "admin",   // 签发人
			IssuedAt:  time.Now().Unix(),
			// 签发时间
		},
		uid,
	}
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//return newWithClaims.SignedString([]byte(global.GvaConfig.Jwt.Secret))
	return newWithClaims.SignedString([]byte(""))
}
