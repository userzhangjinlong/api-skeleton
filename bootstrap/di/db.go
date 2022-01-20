package di

import (
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := ConnectPoolFactory.GetMysql()
	if err != nil {
		panic("db链接获取异常")
	}

	return db
}
