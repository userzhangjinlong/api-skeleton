package ConnectPoolFactory

import "github.com/jinzhu/gorm"

//NewMysql sql工厂加载mysql连接池
func NewMysql() (resule bool) {
	connectErr := NewConnect("mysql").GetInstance().InitConnectPool()

	return connectErr
}

//GetMysql sql工厂获取mysql连接池
func GetMysql() (dbPool *gorm.DB, err error) {
	if !NewMysql() {
		//todo::异常日志记录抛出异常
		panic("mysql连接池获取失败")
	}
	connect, errConnect := NewConnect("mysql").GetInstance().GetConnectLibrary()

	return connect.(*gorm.DB), errConnect
}
