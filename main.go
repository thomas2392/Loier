package main

import (
	redis "loier/config/redis"
	"loier/config/server"
)

func main() {
	redis.CreateRedisClient()
	server.StartServer()
}
