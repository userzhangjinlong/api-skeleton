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

func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	m.UpdateTime = time.Now().Unix()

	return nil
}

//Paginate 自定义分页
func (m *Model) Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
