package ConnectPoolFactory

import (
	"api-skeleton/app/ConstDir"
	"fmt"
	"github.com/go-redis/redis"
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
		//source := getDbLibrary(this.library)
		db, errDb = gorm.Open(mysql.Open(getDbLibrary(ConstDir.DEFAULT)), &gorm.Config{})
		if errDb != nil {
			log.Fatal(errDb.Error())
			return false
		}

		//链接池配置、集群数据源链接配置
		MaxIdleConns, _ := strconv.Atoi(configs.Database.MaxIdleConns)
		MaxOpenConns, _ := strconv.Atoi(configs.Database.MaxOpenConns)
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
		// SetConnMaxLifetime 设置了连接可复用的最大时间
		db.Use(
			dbresolver.
				Register(dbresolver.Config{
					Replicas: []gorm.Dialector{mysql.Open(getDbLibrary(ConstDir.DEFAULT_READ))}, //默认库 读写分离读库
					// sources/replicas 负载均衡策略
					Policy: dbresolver.RandomPolicy{},
				}).
				Register(dbresolver.Config{
					Sources:  []gorm.Dialector{mysql.Open(getDbLibrary(ConstDir.SCHEMA))},      //主库 读写分离写库
					Replicas: []gorm.Dialector{mysql.Open(getDbLibrary(ConstDir.SCHEMA_READ))}, //从库 读写分离读库
					// sources/replicas 负载均衡策略
					Policy: dbresolver.RandomPolicy{},
				}, ConstDir.SCHEMA).
				Register(dbresolver.Config{
					Sources:  []gorm.Dialector{mysql.Open(getDbLibrary(ConstDir.IM))},      //主库 读写分离写库
					Replicas: []gorm.Dialector{mysql.Open(getDbLibrary(ConstDir.IM_READ))}, //从库 读写分离读库
					// sources/replicas 负载均衡策略
					Policy: dbresolver.RandomPolicy{},
				}, ConstDir.IM).
				SetMaxIdleConns(MaxIdleConns).
				SetMaxOpenConns(MaxOpenConns).
				SetConnMaxLifetime(24 * time.Hour).
				SetConnMaxIdleTime(time.Hour),
		)
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
	return &ConnectPool{
		library: library,
	}
}

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
			configs.Database.UserNameIm,
			configs.Database.PasswordIm,
			configs.Database.HostIm,
			configs.Database.PortIm,
			configs.Database.NameIm),
	}
	source := sourceMap[library]

	source += "?charset=" + configs.Database.Charset +
		"&parseTime=True&loc=Local&timeout=5000ms"
	return source
}
