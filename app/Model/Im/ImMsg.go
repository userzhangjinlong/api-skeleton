package Im

import "api-skeleton/app/Model"

type ImMsg struct {

	//数据编号
	Id int64 `gorm:"column:id" json:"id"`

	//消息来源人ID
	FromUserId int64 `gorm:"column:fromUserId" json:"fromUserId"`

	//发送到用户ID
	ToUserId int64 `gorm:"column:toUserId" json:"toUserId"`

	//消息文本内容
	Content string `gorm:"column:content" json:"content"`

	//消息图片内容
	MsgImg string `gorm:"column:msgImg" json:"msgImg"`

	//发送时间
	SendTime int64 `gorm:"column:sendTime" json:"sendTime"`

	//继承父类model
	Model.Model
}

func (model *ImMsg) TableName() string {
	return "im_msg"
}
