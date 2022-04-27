package main

import (
	"github.com/julianfbeck/universal/backend/services/stock/control/route"
	"github.com/julianfbeck/universal/backend/services/stock/shared/repository/redis"
)

func main() {
	redis, err := redis.NewRedisRepository()
	if err != nil {
		panic(err)
	}
	redis.Publish("beju", "hello")
	engine := route.Create()
	engine.Listen(":8080")
}
