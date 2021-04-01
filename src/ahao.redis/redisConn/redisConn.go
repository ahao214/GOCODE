package redisConn

import (
	"ahao.redis/config"
	"github.com/go-redis/redis"
	"log"
)

func ConnectRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	if _, err := conn.Ping().Result(); err != nil {
		log.Fatalf("Connect to redis client failed,err:%v\n", err)
	}
	return conn
}
