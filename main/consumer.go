// main/consumer.go
package main

import (
	"gomall/config"
	"gomall/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("警告：未能从 .env 文件加载环境变量")
	}
	mq := config.InitRabbitMQ()

	// 单独启动消费者
	utils.StartOrderConsumer(mq)
}
