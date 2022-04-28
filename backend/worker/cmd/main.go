package main

import (
	"fmt"

	"github.com/julianfbeck/universal/backend/services/stock/shared/helpers"
	"github.com/julianfbeck/universal/backend/services/stock/shared/repository/redis"
)

type PubMessage struct {
	Response string
}

//main function
func main() {
	redis, err := redis.NewRedisRepository()
	if err != nil {
		panic(err)
	}

	//subscribe channel
	response := make(chan string)
	var repository = PubMessage{}
	go redis.Subscribe("beju", response)

	for {
		err := helpers.StringToType(<-response, &repository)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(repository)
	}
}
