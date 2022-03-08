package Global

import (
	"api-skeleton/config"
	"github.com/go-redis/redis"
	"github.com/nsqio/go-nsq"
	"gorm.io/gorm"
)

var (
	RedisClient  *redis.Client
	RedisCluster *redis.ClusterClient
	DB           *gorm.DB
	Configs      *config.System
	NsqProducer  *nsq.Producer
)
