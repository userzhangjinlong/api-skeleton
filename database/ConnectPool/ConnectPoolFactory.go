package ConnectPoolFactory

import (
	"api-skeleton/app/ConstDir"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	once        sync.Once
	instance    *ConnectPool
	errDb       error
	db          *gorm.DB
	pool        *redis.Client
	poolCluster *redis.ClusterClient
	redisDb     int
	dbType      string
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
		db, errDb = gorm.Open(
			mysql.Open(source), &gorm.Config{})
		if errDb != nil {
			log.Fatal(errDb.Error())
			return false
		}
		//todo::读写分离目前看只能新定义读结构体链接，多读链接如何处理后期学习优化
		//链接池配置、集群数据源链接配置
		MaxIdleConns, _ := strconv.Atoi(configs.Database.MaxIdleConns)
		MaxOpenConns, _ := strconv.Atoi(configs.Database.MaxOpenConns)
		//增加sql配置
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(MaxOpenConns)
		// 设置最大空闲数
		sqlDB.SetMaxIdleConns(MaxIdleConns)
	case "redis":
		redisDb, _ := strconv.Atoi(configs.Redis.Db)
		pool = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", configs.Redis.Root, configs.Redis.Port), // redis地址
			Password: configs.Redis.Auth,                                           // redis密码，没有则留空
			DB:       redisDb,                                                      // 默认数据库，默认是0
		})

		//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
		_, err := pool.Ping().Result()
		if err != nil {
			log.Fatalf("redis链接异常：%s", err)
			return false
		}
	case "redisCluster":
		//redis-cluster 启用方式
		poolCluster = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{
				fmt.Sprintf("%s:%s", configs.RedisCluster.Root, configs.RedisCluster.PortOne),
				fmt.Sprintf("%s:%s", configs.RedisCluster.Root, configs.RedisCluster.PortTwo),
				fmt.Sprintf("%s:%s", configs.RedisCluster.Root, configs.RedisCluster.PortThree),
				fmt.Sprintf("%s:%s", configs.RedisCluster.Root, configs.RedisCluster.PortFour),
				fmt.Sprintf("%s:%s", configs.RedisCluster.Root, configs.RedisCluster.PortFive),
				fmt.Sprintf("%s:%s", configs.RedisCluster.Root, configs.RedisCluster.PortSix),
			},
			Password:     configs.RedisCluster.Auth,
			DialTimeout:  500 * time.Millisecond, // 设置连接超时
			ReadTimeout:  500 * time.Millisecond, // 设置读取超时
			WriteTimeout: 500 * time.Millisecond, // 设置写入超时
		})
		_, err := poolCluster.Ping().Result()
		if err != nil {
			log.Fatalf("redis集群链接异常：%s", err)
			return false
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
	case "redisCluster":
		return poolCluster, err
	default:
		return db, err
	}
}

func NewConnect(connect string, library string) *ConnectPool {
	dbType = connect
	instance = &ConnectPool{
		library: library,
	}

	return instance
}

//getDbLibrary 获取db dsn
func getDbLibrary(library string) string {
	sourceMap := map[string]string{
		ConstDir.DEFAULT: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.Username,
			configs.Database.Password,
			configs.Database.Host,
			configs.Database.Port,
			configs.Database.Name),
		ConstDir.DEFAULT_READ: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.Username,
			configs.Database.Password,
			configs.Database.Host,
			configs.Database.Port,
			configs.Database.Name),
		ConstDir.SCHEMA: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.UsernameSchema,
			configs.Database.PasswordSchema,
			configs.Database.HostSchema,
			configs.Database.PortSchema,
			configs.Database.NameSchema),
		ConstDir.SCHEMA_READ: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.UsernameSchema,
			configs.Database.PasswordSchema,
			configs.Database.HostSchema,
			configs.Database.PortSchema,
			configs.Database.NameSchema),
		ConstDir.IM: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.UserNameIm,
			configs.Database.PasswordIm,
			configs.Database.HostIm,
			configs.Database.PortIm,
			configs.Database.NameIm),

		ConstDir.IM_READ: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configs.Database.UserNameImRead,
			configs.Database.PasswordImRead,
			configs.Database.HostImRead,
			configs.Database.PortImRead,
			configs.Database.NameImRead),
	}
	source := sourceMap[library]

	source += "?charset=" + configs.Database.Charset +
		"&parseTime=True&loc=Local&timeout=5000ms"
	return source
}
