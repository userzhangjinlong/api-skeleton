package Global

import (
	"api-skeleton/config"
	"github.com/garyburd/redigo/redis"
	"gorm.io/gorm"
)

var (
	RedisClient redis.Conn
	DB          *gorm.DB
	Configs     *config.System
)
