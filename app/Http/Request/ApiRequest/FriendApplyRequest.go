package ApiRequest

type ApplyForm struct {
	FriendUid int64  `form:"friendUid" binding:"required,min=0"`
	Desc      string `form:"desc" binding:"required,max=200"`
}

type DealApplyForm struct {
	ApplyId int `form:"applyId" binding:"required"`
	Status  int `form:"status" binding:"oneof=1 2"`
}
