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

func InitRedis() *redis.Client {
	_ = godotenv.Load()

	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")

	dbNum, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Printf("⚠️  REDIS_DB 解析失败，将默认使用 DB 0: %v", err)
		dbNum = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNum,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("❌ 连接 Redis 失败：%v", err)
	}

	fmt.Println("✅ Redis 连接成功")
	return rdb
}

// var RDB *redis.Client
// func InitRedis() {
// 	_ = godotenv.Load() // 加载 .env
// 	addr := os.Getenv("REDIS_ADDR")
// 	password := os.Getenv("REDIS_PASSWORD")
// 	dbStr := os.Getenv("REDIS_DB")
// 	// Redis 客户端需要传 int 类型的数据库编号，但 os.Getenv 返回的是 string
// 	db, err := strconv.Atoi(dbStr)
// 	if  err != nil {
// 		log.Fatalf("REDIS_DB 解析失败，请确认为整数：%v", err)
// 	}

// 	RDB = redis.NewClient(&redis.Options{
// 		Addr:     addr, // 看你本地 redis 配置
// 		Password: password,
// 		DB:       db,
// 	})
// 	if _, err := RDB.Ping(context.Background()).Result(); err != nil {
// 		log.Fatalf("连接 Redis 失败：%v", err)
// 	}
// 	fmt.Println("连接 Redis 成功")
// }
