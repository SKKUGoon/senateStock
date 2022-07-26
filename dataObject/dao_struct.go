package dataObject

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type DataBaseConfig struct {
	Address  string `json:"db"`
	Password string `json:"pw"`
}

type RedisConn[T any] struct {
	Context    context.Context
	Connection *redis.Client
	Container  *T
}

type RdbJson[T any] interface {
	JsonGet(key, path string) error
	JsonSet(key, path string) error
}
