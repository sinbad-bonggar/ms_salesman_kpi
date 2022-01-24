package infra

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type Producer interface {
	SendMessage(topic, msg string) error
}

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducer(
	producer sarama.SyncProducer,
) *KafkaProducer {
	return &KafkaProducer{
		Producer: producer,
	}
}

func (p *KafkaProducer) SendMessage(topic, msg string) error {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		logrus.Errorf("Send message error: %v", err)
		return err
	}

	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}
