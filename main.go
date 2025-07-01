package main

import (
	"gomall/config"
	"gomall/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("警告：未能从 .env 文件加载环境变量")
	}
	config.InitDB()

	r := gin.Default()
	routers.RegisterRoutes(r)

	r.Run(":8080") // 监听端口
}
