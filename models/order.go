package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `gorm:"not null"`
	Status     string  `gorm:"type:varchar(20);default:'pending'"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   gorm.DeletedAt `gorm:"index"` //开启软删除
	OrderItems []OrderItem    `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	Quantity   uint     `gorm:"not null"`
	UnitPrice  float64 `gorm:"type:decimal(10,2);not null"` //下单时的价格，防止价格变动导致混乱
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`

	Product   Product `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"` //开启软删除
}
