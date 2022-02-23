package Util

import (
	"api-skeleton/app/Global"
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
	return newWithClaims.SignedString([]byte(Global.Configs.Jwt.Secret))
}

//ParseToekn 解析token
func ParseToken(tokenString string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Global.Configs.Jwt.Secret), nil
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
