package Api

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request"
	"api-skeleton/app/Model"
	"api-skeleton/app/Util"
	"fmt"
	"github.com/gin-gonic/gin"
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
		return
	}

	//数据表查询用户不存在则创建用户
	var userModel Model.User
	err := Global.DB.Where("tel = ? and password = ?", param.Username, param.Password).Find(&userModel).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userModel)
}
