package ConnectPoolFactory

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
	"sync"
)

type Pool interface {
	GetInstance() *ConnectPool
	InitConnectPool() bool
	GetConnectLibrary() (interface{}, error)
}

var (
	once     sync.Once
	instance *ConnectPool
	errDb    error
	db       *gorm.DB
	pool     *redis.Pool
	redisDb  int
	dbType   string
)

type ConnectPool struct {
	library string
}

func (this *ConnectPool) GetInstance() *ConnectPool {
	once.Do(func() {
		instance = &ConnectPool{
			library: this.library,
		}
	})

	return instance
}

func (this *ConnectPool) InitConnectPool() (result bool) {
	var configs = config.InitConfig
	switch dbType {
	case "mysql":
		source := getDbLibrary(this.library)
		db, errDb = gorm.Open("mysql", source)
		if errDb != nil {
			log.Fatal(errDb.Error())
			return false
		}
		fmt.Println("ConnectPool链接db：", db)
		//关闭数据库连接，db会自动被多个goroutine共享，可以不调用 db貌似不能关闭需要保持长链接？？
		//todo::判断处理db是否需要关闭后期优化
		//defer db.Close()
		log.Println("mysql:初始化连接成功")
	case "redis":
		var redisAddress = configs.Redis.Root + ":" + configs.Redis.Port
		redisDb, _ := strconv.Atoi(configs.Redis.Db)
		pool = &redis.Pool{
			MaxIdle:     10000,
			MaxActive:   0,
			IdleTimeout: 300,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", redisAddress, redis.DialPassword(configs.Redis.Auth), redis.DialDatabase(redisDb))
			},
		}

		defer pool.Close()
		log.Println("redis：实例化连接成功")

	}
	return true
}

func (this *ConnectPool) GetConnectLibrary() (res interface{}, err error) {

	switch dbType {
	case "mysql":
		return db, err
	case "redis":
		return pool, err
	default:
		return db, err
	}
}

func NewConnect(connect string, library string) *ConnectPool {
	dbType = connect
	return &ConnectPool{
		library: library,
	}
}

func getDbLibrary(library string) string {
	var source string
	switch library {
	case ConstDir.DEFAULT:
		source = configs.Database.Username +
			":" + configs.Database.Password +
			"@tcp(" + configs.Database.Host + ":" +
			configs.Database.Port +
			")/" + configs.Database.Name
	case ConstDir.SCHEMA:
		source = configs.Database.UsernameSchema +
			":" + configs.Database.PasswordSchema +
			"@tcp(" + configs.Database.HostSchema + ":" +
			configs.Database.PortSchema +
			")/" + configs.Database.NameSchema
	default:
		source = configs.Database.Username +
			":" + configs.Database.Password +
			"@tcp(" + configs.Database.Host + ":" +
			configs.Database.Port +
			")/" + configs.Database.Name
	}

	source += "?charset=" + configs.Database.Charset +
		"&parseTime=True&loc=Local&timeout=1000ms"
	fmt.Println(source)
	return source
}
