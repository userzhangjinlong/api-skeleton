package Im

import "api-skeleton/app/Model"

type FriendApply struct {

	//用户编号
	Id int64 `gorm:"column:id" json:"id"`

	//用户ID
	UserId int64 `gorm:"column:userId" json:"userId"`

	//申请的好友用户ID
	FriendUserId int64 `gorm:"column:friendUserId" json:"friendUserId"`

	//申请描述内容
	Desc string `gorm:"column:desc" json:"desc"`

	//好友关系状态：0待处理 1 通过 2拒绝
	Status int8 `gorm:"column:status" json:"status"`

	//过期时间
	ExpireTime int64 `gorm:"column:expireTime" json:"expireTime"`

	//继承父类model
	Model.Model
}

func (model *FriendApply) TableName() string {
	return "friend_apply"
}
