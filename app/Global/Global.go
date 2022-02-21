package Global

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

var (
	RedisClient redis.Conn
	DB          *gorm.DB
)
