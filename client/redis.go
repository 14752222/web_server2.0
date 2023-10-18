package client

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"web_server_2.0/config"
)

var ctx = context.Background()

func NewRedisClient(env *config.Env) *redis.Client {
	options := &redis.Options{
		//Addr:         fmt.Sprintf("%s:%d", env.Redis.Host, env.Redis.Port),
		//Password:     "", //env.Redis.Password,
		//DB:           env.Redis.Database,
		//PoolSize:     env.Redis.Pool,
		//IdleTimeout:  time.Duration(env.Redis.IdleTimeoutMillis) * time.Second,
		//WriteTimeout: time.Duration(env.Redis.Timeout) * time.Second,
		//ReadTimeout:  time.Duration(env.Redis.Timeout) * time.Second,
		//DialTimeout:  time.Duration(env.Redis.IdleTimeoutMillis) * time.Second,
		//MinIdleConns: env.Redis.MaxIdle,
		Addr:     fmt.Sprintf("%s:%d", "localhost", 16379),
		Password: "",
		DB:       0,
		PoolSize: 10,
	}

	rdb := redis.NewClient(options)
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis 连接失败", err)
	}
	fmt.Println("Redis 连接成功")
	return rdb
}
