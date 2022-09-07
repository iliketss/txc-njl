package jwtdemo

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"txc-njl/server/model"

	"time"
)

// TokenExpireDuration 有效时间
const TokenExpireDuration = time.Hour * 2

// MySecret 密钥
var mySecret = []byte("i like sisi")

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := model.MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "txc",                                      // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	var mc = new(model.MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
