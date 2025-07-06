package utils

import (
	"encoding/json"
	"gomall/models"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 启动订单消费者
func StartOrderConsumer(conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ 创建 Channel 失败: %v", err)
	}
	defer ch.Close()

	// 声明队列（幂等，确保存在）
	q, err := ch.QueueDeclare(
		"order.created", // 队列名要和生产者一致！
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ 声明队列失败: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // auto-ack
		false, // not exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ 注册消费者失败: %v", err)
	}

	log.Println("📡 订单消费者已启动，等待新订单消息...")

	go func() {
		for msg := range msgs {
			var order models.Order
			if err := json.Unmarshal(msg.Body, &order); err != nil {
				log.Printf("⚠️ 解析订单消息失败: %v", err)
				continue
			}
			// 👉 这里可以替换成你想做的操作，比如发邮件、发通知等
			log.Printf("📥 收到新订单：ID=%d, 用户ID=%d, 总价=%.2f", order.ID, order.UserID, order.TotalPrice)
		}
	}()
}
