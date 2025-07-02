package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	Status     string         `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalPrice float64        `gorm:"type:decimal(10,2);not null" json:"total_price"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeleteAt   gorm.DeletedAt `gorm:"index" json:"-"`
	OrderItems []OrderItem    `gorm:"foreignKey:OrderID" json:"order_items"`
}

const (
	OrderStatusPending  = "pending"  // 等待支付
	OrderStatusPaid     = "paid"     // 已支付
	OrderStatusCanceled = "canceled" // 取消支付
)

type OrderItem struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	Quantity   uint    `gorm:"not null"`
	UnitPrice  float64 `gorm:"type:decimal(10,2);not null"` //下单时的价格，防止价格变动导致混乱
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`

	Product   Product `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"` //开启软删除
}
