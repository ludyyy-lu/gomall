package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	_ = godotenv.Load() // 加载 .env
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")
	// Redis 客户端需要传 int 类型的数据库编号，但 os.Getenv 返回的是 string
	db, _ := strconv.Atoi(dbStr)
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr, // 看你本地 redis 配置
		Password: password,
		DB:       db,
	})
	if _, err := RDB.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("连接 Redis 失败：%v", err)
	}
	fmt.Println("连接 Redis 成功")
}
