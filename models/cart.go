package models

import (
	"time"
)

type CartItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`     // 谁的购物车
	ProductID uint      `gorm:"not null" json:"product_id"`  // 哪个商品
	Quantity  uint      `gorm:"not null" json:"quantity"`    // 购买数量

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"` // 预加载商品信息
}
