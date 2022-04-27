package main

import (
	"fmt"

	"github.com/julianfbeck/universal/backend/services/stock/shared/repository/redis"
)

type RedisRepository struct {
	HI string
}

//main function
func main() {
	redis, err := redis.NewRedisRepository()
	if err != nil {
		panic(err)
	}

	//subscribe channel
	response := make(chan string)
	// var repository = RedisRepository{}
	go redis.Subscribe("beju", response)

	for {
		fmt.Println(<-response)
		// helpers.StringToType(<-response, &repository)
	}
}
