package main

import (
	"github.com/julianfbeck/universal/backend/services/stock/control/route"
	"github.com/julianfbeck/universal/backend/services/stock/shared/repository/redis"
)

type PubMessage struct {
	Response string
}

func main() {
	redis, err := redis.NewRedisRepository()
	if err != nil {
		panic(err)
	}
	//send PubMessage to redis
	message := PubMessage{
		Response: "hello",
	}
	err = redis.Publish("beju", message)
	if err != nil {
		panic(err)
	}
	engine := route.Create()
	engine.Listen(":8080")
}
