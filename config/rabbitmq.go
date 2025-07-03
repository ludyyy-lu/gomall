package config

import (
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ() *amqp091.Connection {
	url := os.Getenv("RABBITMQ_URL")
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("连接 RabbitMQ 失败：%v", err)
		return nil
	}
	fmt.Println("✅ RabbitMQ 连接成功")
	return conn
}
