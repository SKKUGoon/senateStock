package dataObject

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func processKey(keyAddr string) (DataBaseConfig, error) {
	var c DataBaseConfig
	file, err := os.Open(keyAddr)
	if err != nil {
		return c, err
	}
	byteVal, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteVal, &c)
	return c, err
}

func Conn(keyAddr string) (*redis.Client, error) {
	loginInfo, err := processKey(keyAddr)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	fmt.Println("go redis tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     loginInfo.Address,
		Password: loginInfo.Password,
		DB:       0,
	})
	fmt.Println("connection successful")
	return client, nil
}
