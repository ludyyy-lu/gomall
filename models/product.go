package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Price       float64    `gorm:"not null" json:"price"`
	Stock       uint       `gorm:"default:0" json:"stock"`
	ImageURL    string     `gorm:"type:varchar(255)" json:"image_url"`
	UserID      uint       `gorm:"not null" json:"user_id"`      // 创建者
	Categories  []Category `gorm:"many2many:product_categories"` // 多对多
	OnSale      bool       `gorm:"default:true" json:"on_sale"`
	Version     int        `gorm:"default:1" json:"version"` // 乐观锁版本号
	IsSeckill   bool       `gorm:"default:false" json:"is_seckill"` // 是否秒杀商品
}
