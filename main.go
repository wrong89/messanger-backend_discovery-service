package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://:123@localhost:6380/0")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	data, err := client.Ping(context.TODO()).Result()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Discovery service is started on localhost:9091")

	fmt.Println("Successfully ping", data)
	http.ListenAndServe(":9091", nil)
}
