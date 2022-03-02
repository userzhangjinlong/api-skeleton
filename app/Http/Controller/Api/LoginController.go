package Api

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request"
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/app/Util"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Login struct {
}

func (l *Login) Login(ctx *gin.Context) {
	//参数校验
	param := struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required,min=6"`
	}{}
	valid, errs := Request.BindAndValid(ctx, &param)
	if !valid {
		Util.Error(ctx, 100, fmt.Sprintf("参数错误：%s", errs))
	}

	//数据表查询用户不存在则创建用户
	var userModel ApiSkeleton.User
	err := Global.DB.Where("tel = ? and password = ?", param.Username, param.Password).Find(&userModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//数据不存在新增数据
		userModel.Tel = param.Username
		userModel.Username = param.Username
		userModel.Password = param.Password
		userModel.CreateTime = time.Now().Unix()
		userModel.UpdateTime = time.Now().Unix()
		err = Global.DB.Create(&userModel).Error
		if err != nil {
			Util.Error(ctx, 100, fmt.Sprintf("登陆失败：%s", err))
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
	Util.Success(ctx, token)
}
