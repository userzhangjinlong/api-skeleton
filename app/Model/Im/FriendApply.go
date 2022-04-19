package Im

import "api-skeleton/app/Model"

type FriendApply struct {
   
	//用户编号  
	 Id int64 `gorm:"id" json:"id"` 
  
	//用户ID  
	 UserId int64 `gorm:"userId" json:"userId"` 
  
	//申请的好友用户ID  
	 FriendUserId int64 `gorm:"friendUserId" json:"friendUserId"` 
  
	//申请描述内容  
	 Desc string `gorm:"desc" json:"desc"` 
  
	//好友关系状态：0待处理 1 通过 2拒绝  
	 Status int8 `gorm:"status" json:"status"` 
  
	//过期时间  
	 ExpireTime int64 `gorm:"expireTime" json:"expireTime"` 

	//继承父类model
	Model.Model
}

func (model *FriendApply) TableName() string {
	return "friend_apply"
}