package Global

import (
	"api-skeleton/config"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

var (
	RedisClient redis.Conn
	DB          *gorm.DB
	Configs     *config.System
)
