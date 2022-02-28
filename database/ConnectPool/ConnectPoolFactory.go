package ConnectPoolFactory

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Model/InformationSchema"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
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
		//source := getDbLibrary(this.library)
		db, errDb = gorm.Open(mysql.Open(getDbLibrary(ConstDir.DEFAULT)), &gorm.Config{})
		if errDb != nil {
			log.Fatal(errDb.Error())
			return false
		}

		//链接池配置
		MaxIdleConns, _ := strconv.Atoi(configs.Database.MaxIdleConns)
		MaxOpenConns, _ := strconv.Atoi(configs.Database.MaxOpenConns)
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
		// SetConnMaxLifetime 设置了连接可复用的最大时间
		db.Use(
			dbresolver.Register(dbresolver.Config{
				Sources:  []gorm.Dialector{mysql.Open(getDbLibrary(ConstDir.SCHEMA))},
				Replicas: []gorm.Dialector{mysql.Open(getDbLibrary(ConstDir.SCHEMA))},
				// sources/replicas 负载均衡策略
				Policy: dbresolver.RandomPolicy{},
			}, &InformationSchema.Columns{}).
				SetMaxIdleConns(MaxIdleConns).
				SetMaxOpenConns(MaxOpenConns).
				SetConnMaxLifetime(24 * time.Hour).
				SetConnMaxIdleTime(time.Hour),
		)
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
		source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.Username,
			configs.Database.Password,
			configs.Database.Host,
			configs.Database.Port,
			configs.Database.Name)
	case ConstDir.SCHEMA:
		source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.UsernameSchema,
			configs.Database.PasswordSchema,
			configs.Database.HostSchema,
			configs.Database.PortSchema,
			configs.Database.NameSchema)
	default:
		source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.Username,
			configs.Database.Password,
			configs.Database.Host,
			configs.Database.Port,
			configs.Database.Name)
	}

	source += "?charset=" + configs.Database.Charset +
		"&parseTime=True&loc=Local&timeout=1000ms"
	return source
}
