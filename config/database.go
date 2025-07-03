package config

import (
	"fmt"
	"gomall/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

// func InitDB() {
// 	_ = godotenv.Load() // 加载 .env

// 	dsn := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASS"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("MySQL数据库连接失败:", err)
// 	}

// 	fmt.Println("✅ MySQL数据库连接成功")
// 	err = DB.AutoMigrate(
// 		&models.User{},
// 		&models.Product{},
// 		&models.CartItem{},
// 		&models.Order{},
// 		&models.OrderItem{},
// 	)
// 	if err != nil {
// 		log.Fatal("❌ MySQL自动迁移失败:", err)
// 	}
// 	fmt.Println("✅ MySQL模型迁移成功")
// }

// InitDBWithReturn 初始化并返回 DB 实例
func InitDB() (*gorm.DB, error) {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("MySQL连接数据库失败: %w", err)
	}
	fmt.Println("✅ MySQL数据库连接成功")
	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("❌ MySQL自动迁移失败:", err)
	}
	fmt.Println("✅ MySQL模型迁移成功")
	return db, nil
}
