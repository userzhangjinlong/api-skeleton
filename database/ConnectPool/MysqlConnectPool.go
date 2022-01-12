package ConnectPoolFactory

import "github.com/jinzhu/gorm"

//mysql工厂加载mysql连接池
func NewMysql() (resule bool) {
	connectErr := NewConnect("mysql").GetInstance().InitConnectPool()

	return connectErr
}

//mysql工厂获取mysql连接池
func GetMysql() (dbPool *gorm.DB, err error) {
	connect,errConnect := NewConnect("mysql").GetInstance().GetConnectLibrary()

	return connect.(*gorm.DB),errConnect
}
