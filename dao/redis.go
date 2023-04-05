package dao

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func init() {
	ctx := context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6391",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Error("[redis][init] redis init error: ", err)
		return
	}
}

func Set(key string, val interface{}, dur int) error {
	rdb.SetNX(context.Background(), key, val, time.Duration(dur))

	return nil
}
