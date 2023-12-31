package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	settings "tieba/setting"
)

// 声明一个全局的rdb变量
var client *redis.Client

//var ctx = context.Background()

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	return err
}
func Close() {
	_ = client.Close()
}
