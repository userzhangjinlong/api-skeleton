package Util

import (
	"api-skeleton/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	ID         string `json:"userId"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

func CreateToken(user *UserClaims) (string, error) {
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	var configs = config.InitConfig
	return newWithClaims.SignedString([]byte(configs.Jwt.Secret))
}

//ParseToekn 解析token
func ParseToken(tokenString string) (*UserClaims, error) {
	var configs = config.InitConfig

	token, err := jwt.ParseWithClaims(tokenString, UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	userClaims, ok := token.Claims.(*UserClaims)

	if !ok || !token.Valid {
		return nil, errors.New("token解析失败")
	}

	return userClaims, nil
}
