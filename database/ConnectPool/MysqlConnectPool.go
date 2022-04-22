package ConnectPoolFactory

import (
	"gorm.io/gorm"
)

//NewMysql sql工厂加载mysql连接池
func NewMysql(library string) (resule bool) {
	connectErr := NewConnect("mysql", library).InitConnectPool()

	return connectErr
}

//GetMysql sql工厂获取mysql连接池
func GetMysql(library string) (dbPool *gorm.DB, err error) {
	if !NewMysql(library) {
		//todo::异常日志记录抛出异常
		panic("mysql连接池获取失败")
	}

	connect, errConnect := NewConnect("mysql", library).GetConnectLibrary()

	return connect.(*gorm.DB), errConnect
}
