package store

import (
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	pinged := rdb.Ping()
	if pinged.Err() == nil {
		log.Println("✅ redis connection successful")
	}
}

func GetRedis() *redis.Client {
	return rdb
}
