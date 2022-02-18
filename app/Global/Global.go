package Global

import (
	"github.com/garyburd/redigo/redis"
)

var (
	RedisClient redis.Conn
)
