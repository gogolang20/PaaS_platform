package dao

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type IRedisDB interface {
	Set(ctx context.Context, key string, value interface{}) error
}

type RedisDB struct {
	redisdb *redis.Client
}

func NewRedisDB() IRedisDB {
	ctx := context.Background()

	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6391",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Error("[redis][NewRedisDB] ping redis error: ", err)
		return nil
	}

	return &RedisDB{redisdb: redisdb}
}

func (rdb *RedisDB) Set(ctx context.Context, key string, value interface{}) error {
	nx := rdb.redisdb.SetNX(ctx, key, value, time.Duration(300))
	result, err := nx.Result()
	if err != nil {
		return err
	}
	if result {
		return nil
	}

	return nil
}

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
