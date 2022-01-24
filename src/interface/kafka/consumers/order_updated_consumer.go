package consumers

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type orderUpdatedConsumer struct{}

func NewOrderUpdatedConsumer() ConsumerInterface {
	return &orderUpdatedConsumer{}
}

var _ ConsumerInterface = &orderUpdatedConsumer{}

func (k *orderUpdatedConsumer) Consume(m *sarama.ConsumerMessage) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover")
		}
	}()

	logrus.Debugf("New message consumed from %v", m.Partition)
	logrus.Info(string(m.Value))
	return nil
}
