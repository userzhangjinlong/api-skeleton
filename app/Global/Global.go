package Global

import (
	"api-skeleton/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	RedisClient  *redis.Client
	RedisCluster *redis.ClusterClient
	DB           *gorm.DB
	Configs      *config.System
)
