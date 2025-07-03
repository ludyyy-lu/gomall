package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MQConn *amqp.Connection

func InitRabbitMQ() {
	_ = godotenv.Load()
	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("连接 RabbitMQ 失败：%v", err)
	}
	MQConn = conn
}
