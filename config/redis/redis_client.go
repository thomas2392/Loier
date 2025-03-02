package redis_client

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func CreateRedisClient() {
	var ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Error when establishing redis connection.")
		log.Fatal(err)
	}
	log.Println("Conection with redis sucessfull.")
}

func GetRedisClient() *redis.Client {
	return rdb
}
