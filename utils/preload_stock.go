package utils

import (
	"context"
	"fmt"
	"gomall/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func PreloadSeckillStock(db *gorm.DB, rdb *redis.Client) error {
	var products []models.Product

	// 查找参与秒杀的商品，例如你可以有个字段 is_seckill
	if err := db.Where("on_sale = ? AND stock > 0 AND is_seckill = ?", true, true).
	Find(&products).Error; err != nil {
		return err
	}

	for _, product := range products {
		key := fmt.Sprintf("seckill:stock:%d", product.ID)
		err := rdb.Set(context.Background(), key, product.Stock, 0).Err()
		if err != nil {
			return fmt.Errorf("写入 Redis 失败，商品 ID=%d: %w", product.ID, err)
		}
	}

	fmt.Println("✅ 秒杀商品库存已预加载到 Redis")
	return nil
}
