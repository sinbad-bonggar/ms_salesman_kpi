package consumers

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type orderCreatedConsumer struct{}

func NewOrderCreatedConsumer() ConsumerInterface {
	return &orderCreatedConsumer{}
}

var _ ConsumerInterface = &orderCreatedConsumer{}

func (k *orderCreatedConsumer) Consume(m *sarama.ConsumerMessage) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover")
		}
	}()

	logrus.Infof("New message consumed from Partition %d", m.Partition)
	logrus.Info(string(m.Value))
	return nil
}
