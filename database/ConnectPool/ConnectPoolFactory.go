package ConnectPoolFactory

import (
	"api-skeleton/app/ConstDir"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
	"sync"
	"time"
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
	switch dbType {
	case "mysql":
		source := getDbLibrary(this.library)
		db, errDb = gorm.Open("mysql", source)
		if errDb != nil {
			log.Fatal(errDb.Error())
			return false
		}
		db.SingularTable(true)
		//链接池配置
		MaxIdleConns, _ := strconv.Atoi(configs.Database.MaxIdleConns)
		MaxOpenConns, _ := strconv.Atoi(configs.Database.MaxOpenConns)
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
		db.DB().SetMaxIdleConns(MaxIdleConns)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		db.DB().SetMaxOpenConns(MaxOpenConns)
		// SetConnMaxLifetime 设置了连接可复用的最大时间
		db.DB().SetConnMaxLifetime(time.Hour)
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
