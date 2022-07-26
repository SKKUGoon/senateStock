package dataObject

import (
	"encoding/json"
	"log"
)

func (rdb RedisConn[any]) JsonGet(key, path string) error {
	log.Printf("JSON GET: %v %v Called\n", key, path)
	cmd := rdb.Connection.Do(rdb.Context, "JSON.GET", key, path)
	v, _ := cmd.Result()

	if vs, ok := v.(string); ok {
		err := json.Unmarshal([]byte(vs), &rdb.Container)
		if err != nil {
			log.Panicln(err)
			return err
		}
	}
	return nil
}

func (rdb RedisConn[any]) JsonSet(key, path string) error {
	setVal, _ := json.Marshal(rdb.Container)
	log.Printf("JSON SET : %v:: %v Setting", key, path)
	cmd := rdb.Connection.Do(rdb.Context, "JSON.SET", key, path, string(setVal))
	err := cmd.Err()
	return err
}
