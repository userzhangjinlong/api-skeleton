package ApiSkeleton

type ApiDoubleFlowBringRankRealtimeIncr1d struct {

	//自增ID  
	Id int64 `gorm:"id" json:"id"`

	//店铺id  
	ShopId int64 `gorm:"shop_id" json:"shop_id"`

	//主营一级类目ID  
	FirstCategoryId int64 `gorm:"first_category_id" json:"first_category_id"`

	//人气值  
	PopValue float64 `gorm:"pop_value" json:"pop_value"`

	//排名  
	Ranking int64 `gorm:"ranking" json:"ranking"`

	//维护字段-更新时间  
	ModifyTime string `gorm:"modify_time" json:"modify_time"`

	//分区日期  
	Dt string `gorm:"dt" json:"dt"`
}

func (model *ApiDoubleFlowBringRankRealtimeIncr1d) TableName() string {
	return "api_double_flow_bring_rank_realtime_incr_1d"
}