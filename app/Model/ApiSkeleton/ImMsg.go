package ApiSkeleton

type ImMsg struct {
   
	//数据编号  
	 Id int64 `gorm:"id" json:"id"` 
  
	//消息来源人ID  
	 FromUserId int64 `gorm:"fromUserId" json:"fromUserId"` 
  
	//发送到用户ID  
	 ToUserId int64 `gorm:"toUserId" json:"toUserId"` 
  
	//消息文本内容  
	 Content string `gorm:"content" json:"content"` 
  
	//消息图片内容  
	 MsgImg string `gorm:"msgImg" json:"msgImg"` 
  
	//发送时间  
	 SendTime int64 `gorm:"sendTime" json:"sendTime"` 
  
	//创建时间  
	 CreateTime int64 `gorm:"createTime" json:"createTime"` 
  
	//更新时间  
	 UpdateTime int64 `gorm:"updateTime" json:"updateTime"` 

}

func (model *ImMsg) TableName() string {
	return "im_msg"
}