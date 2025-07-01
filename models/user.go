package models
import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"size:64;uniqueIndex" json:"username"`
    Email    string `gorm:"size:128;uniqueIndex" json:"email"`
    Password string `json:"-"` // 不返回密码
}