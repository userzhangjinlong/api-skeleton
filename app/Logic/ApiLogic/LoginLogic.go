package ApiLogic

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request/ApiRequest"
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/app/Util"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func LoginLogic(form *ApiRequest.LoginForm) (string, error) {
	//数据表查询用户不存在则创建用户
	var userModel ApiSkeleton.User
	err := Global.DB.
		Where("tel = ? and password = ?", form.Username, Util.Md5Encryption(form.Password)).
		Find(&userModel).Error
	if err == gorm.ErrRecordNotFound {
		//数据不存在新增数据
		userModel.Tel = form.Username
		userModel.Username = form.Username
		userModel.Password = Util.Md5Encryption(form.Password)
		err = Global.DB.Create(&userModel).Error
		if err != nil {
			return "", err
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

	return token, nil
}
