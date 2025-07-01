package models

import (
    "gorm.io/gorm"
)

type Category struct {
    ID          uint           `gorm:"primaryKey"`
    Name        string         `gorm:"type:varchar(100);not null;unique"`
    Description string         `gorm:"type:text"`
    gorm.DeletedAt `gorm:"index"`
    Products    []Product      `gorm:"many2many:product_categories"`
}
