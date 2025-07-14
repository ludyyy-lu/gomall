package utils

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// 扣减秒杀库存的 Lua 脚本内容（也可以抽出来放到独立文件）
const luaScript = `
local stock = tonumber(redis.call("GET", KEYS[1]))
if stock == nil then
    return -1
end
if stock <= 0 then
    return 0
end
redis.call("DECRBY", KEYS[1], ARGV[1])
return 1
`

func SeckillDecrStock(rdb *redis.Client, productID uint, quantity uint) (int64, error) {
	key := fmt.Sprintf("seckill:stock:%d", productID)

	result, err := rdb.Eval(context.Background(), luaScript, []string{key}, quantity).Result()
	if err != nil {
		return 0, err
	}

	// 返回值说明：
	// -1: key 不存在
	//  0: 库存不足
	//  1: 成功
	return result.(int64), nil
}
