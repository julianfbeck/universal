package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/julianfbeck/universal/backend/services/stock/shared/helpers"
	"github.com/julianfbeck/universal/backend/services/stock/shared/repository/redis"
)

type PubMessage struct {
	Response string
}

var (
	NUMBER_OF_WORKERS = runtime.NumCPU()
	METRICS_PORT      = "2112"
)

//main function
func main() {
	log.Printf("Starting %d workers...", NUMBER_OF_WORKERS)
	go metrics()
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

func metrics() {
	log.Printf("Starting metrics server on port %s", METRICS_PORT)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":"+METRICS_PORT, nil)
	if err != nil {
		panic(err)
	}
}
