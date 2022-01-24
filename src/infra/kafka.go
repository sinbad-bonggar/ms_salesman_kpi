package infra

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/kafka/consumers"
	"github.com/sirupsen/logrus"
)

func RegisterKafka() (*KafkaConsumer, *KafkaProducer, error) {
	conf := getKafkaConfig("", "")

	prod, err := registerKafkaProducer(conf)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return nil, nil, err
	}

	cons, err := registerKafkaConsumer(conf)
	if err != nil {
		logrus.Errorf("Unable to create kafka consumer got error %v", err)
		return nil, nil, err
	}

	go cons.Consume()
	return cons, prod, nil
}

func registerKafkaProducer(conf *sarama.Config) (*KafkaProducer, error) {
	producers, err := sarama.NewSyncProducer([]string{"localhost:29092"}, conf)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return nil, err
	}
	return &KafkaProducer{Producer: producers}, nil
}

func registerKafkaConsumer(conf *sarama.Config) (*KafkaConsumer, error) {
	client, err := sarama.NewClient([]string{"localhost:29092"}, conf)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return nil, err
	}

	consumers, err := sarama.NewConsumerGroupFromClient("order-consumer", client)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return nil, err
	}

	return &KafkaConsumer{
		Consumer:       consumers,
		ConsumerGroups: getKafkaConsumers(),
		Signals:        make(chan os.Signal, 1),
	}, nil
}

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}

// Here you put all your topics and each consumer
func getKafkaConsumers() map[string]KafkaConsumersConfig {
	return map[string]KafkaConsumersConfig{
		"order.created": {
			Consumer: consumers.NewOrderCreatedConsumer(),
		},
		"order.updated": {
			Consumer: consumers.NewOrderUpdatedConsumer(),
		},
	}
}
