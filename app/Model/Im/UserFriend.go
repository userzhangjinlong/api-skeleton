package Im

import "api-skeleton/app/Model"

const (
	NORMAL_STATUS = iota
	BLACK_STATUS
	DELETE_STATUS
)

type UserFriend struct {

	//用户编号
	Id int64 `gorm:"column:id" json:"id"`

	//用户ID
	UserId int64 `gorm:"column:userId" json:"userId"`

	//好友用户ID
	FriendUserId int64 `gorm:"column:friendUserId" json:"friendUserId"`

	//好友关系状态：0正常，1拉黑，2删除
	Status int8 `gorm:"status" json:"status"`

	//继承父类model
	Model.Model
}

func (model *UserFriend) TableName() string {
	return "user_friend"
}
