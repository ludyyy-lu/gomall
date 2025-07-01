package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"default:0" json:"stock"`
	ImageURL    string  `gorm:"type:varchar(255)" json:"image_url"`
	UserID      uint    `gorm:"not null" json:"user_id"` // 创建者
}
