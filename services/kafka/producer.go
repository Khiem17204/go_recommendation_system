package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func (k *KafkaProducer) NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	producer, err := NewProducer(brokers)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{producer: producer}, nil
}

func NewProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %v", err)
	}

	return producer, nil
}

func (k *KafkaProducer) Close() {
	k.producer.Close()
}

func (k *KafkaProducer) SendMessage(topic string, key string, value string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	_, _, err := k.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	return nil
}
