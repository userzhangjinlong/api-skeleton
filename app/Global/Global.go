package Global

import (
	"api-skeleton/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	RedisClient *redis.Client
	DB          *gorm.DB
	Configs     *config.System
)
