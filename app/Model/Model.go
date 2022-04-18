package Model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	//创建时间
	CreateTime int64 `gorm:"column:createTime" json:"createTime"`
	//更新时间
	UpdateTime int64 `gorm:"column:updateTime" json:"updateTime"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.CreateTime = time.Now().Unix()
	m.UpdateTime = time.Now().Unix()

	return nil
}
