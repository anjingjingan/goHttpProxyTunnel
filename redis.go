package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var Ctx = context.Background()
var Redis *redis.Client

func RedisClient() {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"), // 没有密码，默认值
		DB:       db,                          // 默认DB 0
	})
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		panic("redis连接错误：" + err.Error())
	}

	Redis = rdb
}
