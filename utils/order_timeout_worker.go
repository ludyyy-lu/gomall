package utils

import (
	"context"
	"fmt"
	"gomall/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 后台定时任务定期检查并取消订单
func StartOrderTimeoutWatcher(rdb *redis.Client, db *gorm.DB) {
	go func() {
		for {
			now := time.Now().Unix()
			res, _ := rdb.ZRangeByScore(context.Background(), "order:timeout", &redis.ZRangeBy{
				Min: "-inf",
				Max: fmt.Sprintf("%d", now),
			}).Result()

			for _, orderIDStr := range res {
				orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
				if err != nil {
					continue
				}

				var order models.Order
				if err := db.Preload("OrderItems").First(&order, uint(orderID)).Error; err != nil {
					continue
				}
				if order.Status == "pending" {
					// 设置状态为超时
					order.Status = "timeout"
					// 恢复库存
					if order.OrderItems != nil {
						for _, item := range order.OrderItems {
							var product models.Product
							if err := db.First(&product, item.ProductID).Error; err == nil {
								product.Stock += item.Quantity
								db.Save(&product)
							}
						}
					}

					// 保存订单状态
					db.Save(&order)
				}
				// 从 Redis 移除
				rdb.ZRem(context.Background(), "order:timeout", orderIDStr)
			}

			time.Sleep(1 * time.Minute)
		}
	}()
}
