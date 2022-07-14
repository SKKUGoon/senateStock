package dataObject

import (
	"fmt"
	"github.com/go-redis/redis"
)

func Conn() {
	fmt.Println("go redis tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
