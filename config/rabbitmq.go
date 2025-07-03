package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

var MQConn *amqp091.Connection

func InitRabbitMQ() {
	_ = godotenv.Load()
	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("连接 RabbitMQ 失败：%v", err)
	}
	MQConn = conn
}
