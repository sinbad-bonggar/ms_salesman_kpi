package consumers

import "github.com/Shopify/sarama"

type ConsumerInterface interface {
	Consume(m *sarama.ConsumerMessage) error
}

const (
	OrderCreatedTopic = "order.created"
	OrderUpdatedTopic = "order.updated"
)
