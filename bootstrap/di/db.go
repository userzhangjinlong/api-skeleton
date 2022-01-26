package di

import (
	"api-skeleton/app/ConstDir"
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := ConnectPoolFactory.GetMysql(ConstDir.DEFAULT)
	if err != nil {
		panic("db链接获取异常")
	}

	return db
}
