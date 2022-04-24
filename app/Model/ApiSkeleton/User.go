package ApiSkeleton

import (
	"api-skeleton/app/Model"
	"api-skeleton/app/Util"
)

type User struct {

	//用户编号
	Id int `gorm:"column:id" json:"id"`

	//用户名
	Username string `gorm:"column:username" json:"username"`

	//密码
	Password string `gorm:"column:password" json:"password"`

	//邮箱
	Email string `gorm:"column:email" json:"email"`

	//年龄
	Age int8 `gorm:"column:age" json:"age"`

	//电话
	Tel string `gorm:"column:tel" json:"tel"`

	//地址
	Addr string `gorm:"column:addr" json:"addr"`

	//自定义头像
	Avatar string `gorm:"column:avatar" json:"avatar"`

	//继承父类model
	Model.Model
}

var UserInfo *Util.UserClaims

func (model *User) TableName() string {
	return "user"
}
