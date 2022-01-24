package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/kafka/consumers"
	"github.com/sirupsen/logrus"
)

type KafkaConsumersConfig struct {
	Consumer consumers.ConsumerInterface
}

func (c *KafkaConsumersConfig) GetAllTopics() {

}

type KafkaConsumer struct {
	Consumer       sarama.ConsumerGroup
	ConsumerGroups map[string]KafkaConsumersConfig
	Signals        chan os.Signal
}

func (c *KafkaConsumer) Consume() {

	var topics []string
	for k := range c.ConsumerGroups {
		topics = append(topics, k)
	}

	handler := consumerGroupHandler{
		Signal:         c.Signals,
		ConsumerGroups: c.ConsumerGroups,
	}

	c.Consumer.Consume(context.Background(), topics, handler)

	defer func() {
		if err := c.Consumer.Close(); err != nil {
			logrus.Errorf("Unable to close group", err)
		}
	}()
}

type consumerGroupHandler struct {
	Signal         chan os.Signal
	ConsumerGroups map[string]KafkaConsumersConfig
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	logrus.Infof("Kafka consumer %s is started . . .", claim.Topic())
	go consumeMessage(sess, claim, chanMessage)

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			fmt.Println(string(msg.Value))
			fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)

			h.ConsumerGroups[claim.Topic()].Consumer.Consume(msg)
		case sig := <-h.Signal:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
	return nil
}
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func consumeMessage(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim, c chan *sarama.ConsumerMessage) {
	for {
		msg := <-claim.Messages()
		c <- msg

		sess.MarkMessage(msg, "")
	}
}
