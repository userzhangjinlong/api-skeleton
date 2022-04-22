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
	DB           *gorm.DB //默认db apiSkeleton
	ImDB         *gorm.DB //IMDB
	SchemaDB     *gorm.DB //SchemaDB
	Configs      *config.System
	NsqProducer  *nsq.Producer
	NsqProducer2 *nsq.Producer
	NsqProducer3 *nsq.Producer
)
