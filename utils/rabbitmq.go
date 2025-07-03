package utils

import (
	"encoding/json"
	"gomall/models"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 生产者发送订单消息
// 发送订单创建消息
func PublishOrderCreated(ch *amqp.Channel, order models.Order) error {
	// RabbitMQ 只能传输字节数据，不能直接传 Go 的 struct，所以我们用 json.Marshal 转成 []byte。
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// 确保队列存在（幂等操作）
	// 如果队列不存在就创建，如果已存在就跳过
	_, err = ch.QueueDeclare(
		"order.created", // 队列名
		true,            // durable：是否持久化（true表示重启 RabbitMQ 后队列还在）
		false,           // autoDelete：无消费者时是否自动删除（false 表示保留）
		false,           // exclusive：是否排他（true表示只允许一个连接用，通常是false）
		false,           // noWait：是否异步声明，false=等RabbitMQ响应
		nil,             // args：其他参数（暂时不填）
	)
	if err != nil {
		return err
	}

	// 发送消息
	err = ch.Publish(
		"",              // Exchange 为空，表示使用默认交换机
		"order.created", // RoutingKey 路由到哪个队列
		false,           // mandatory：是否要求消息必须被投递（false = 不强制）
		false,           // immediate：是否要求马上被消费者接收（已废弃，false）
		amqp.Publishing{
			ContentType: "application/json", // 告诉消费者这是 JSON
			Body:        body,               // 实际消息体（[]byte）
		})

	if err != nil {
		return err
	}

	log.Println("✅ 成功发送订单创建消息到队列")
	return nil
}
