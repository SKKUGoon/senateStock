package dataObject

import (
	"log"
)

func (rdb RedisConn[any]) HashGet(key string) (string, error) {
	log.Printf("HASH GET: %v", key)
	val, err := rdb.Connection.Get(rdb.Context, key).Result()
	if err != nil {
		log.Panicln(err)
	}
	return val, err
}

func (rdb RedisConn[any]) HashSet(key, value string) error {
	log.Printf("HASH SET: %v w/ %v", key, value)
	err := rdb.Connection.Set(rdb.Context, key, value, 0).Err()
	if err != nil {
		log.Panicln(err)
	}
	return err
}
