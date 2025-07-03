package utils

import "github.com/redis/go-redis/v9"

var RDB *redis.Client

func SetupRedis(rdb *redis.Client) {
	RDB = rdb
}
