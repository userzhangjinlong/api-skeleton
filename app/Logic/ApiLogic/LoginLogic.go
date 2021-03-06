package ApiLogic

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request/ApiRequest"
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/app/Util"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type LoginData struct {
	Token      string `json:"token"`
	UserId     int    `json:"userId"`
	Username   string `json:"username"`
	Tel        string `json:"tel"`
	Avatar     string `json:"avatar"`
	CreateTime int64  `json:"createTime"`
}

func LoginLogic(form *ApiRequest.LoginForm) (*LoginData, error) {
	//数据表查询用户不存在则创建用户
	var userModel ApiSkeleton.User
	var count int64
	Global.DB.
		Where("tel = ? and password = ?", form.Username, Util.Md5Encryption(form.Password)).
		Find(&userModel).Count(&count)
	if count == 0 {
		//数据不存在新增数据
		userModel.Tel = form.Username
		userModel.Username = form.Username
		userModel.Password = Util.Md5Encryption(form.Password)
		err := Global.DB.Create(&userModel).Error
		if err != nil {
			return nil, err
		}
	}
	expireTime, _ := time.ParseDuration(Global.Configs.Jwt.Expire)
	userClaims := Util.UserClaims{
		ID:         strconv.Itoa(userModel.Id),
		Name:       userModel.Username,
		Phone:      userModel.Tel,
		CreateTime: strconv.Itoa(int(userModel.CreateTime)),
		UpdateTime: strconv.Itoa(int(userModel.UpdateTime)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
			Issuer:    Global.Configs.Jwt.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token, _ := Util.CreateToken(&userClaims)
	LoginData := LoginData{
		Token:      token,
		UserId:     userModel.Id,
		Username:   userModel.Username,
		Tel:        userModel.Tel,
		Avatar:     userModel.Avatar,
		CreateTime: userModel.CreateTime,
	}
	return &LoginData, nil
}
