package common

import (
	"github.com/dgrijalva/jwt-go"
	"idleRain.com/ginEssential/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 生成 token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 一个星期有效期
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // token 有效期
			IssuedAt:  time.Now().Unix(),     // token 发放时间
			Issuer:    "idleRain.com",        // token 发放机构
			Subject:   "user token",          // 主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) // 使用密钥生成 token
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokeString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokeString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
