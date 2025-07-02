package models

import "time"

type Order struct {
	ID         uint         `gorm:"primaryKey"`
	UserID     uint         `gorm:"not null"`
	Status     string       `gorm:"type:varchar(20);default:'pending'"`
	TotalPrice float64      `gorm:"type:decimal(10,2);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	Quantity   int     `gorm:"not null"`
	UnitPrice  float64 `gorm:"type:decimal(10,2);not null"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`
	CreatedAt  time.Time
}
