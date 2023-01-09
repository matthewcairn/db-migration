package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

type event struct {
	Type  string
	Value interface{}
}

// NewKafkaProducer instantiates the Kafka producer using configuration defined in environment variables.
func NewKafkaProducer() (*KafkaProducer, error) {
	host, topic, err := newKafkaConfig()
	if err != nil {
		return nil, fmt.Errorf("newKafkaConfig %w", err)
	}

	config := kafka.ConfigMap{
		"bootstrap.servers": host,
	}

	client, err := kafka.NewProducer(&config)
	if err != nil {
		return nil, fmt.Errorf("kafka.NewProducer %w", err)
	}

	return &KafkaProducer{
		Producer: client,
		Topic:    topic,
	}, nil
}

func (k *KafkaProducer) PublishSimple(Topic string, key string, value string) error {
	if err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &Topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(value),
		Key:   []byte(key),
	}, nil); err != nil {
		return err
	}

	return nil
}

func (k *KafkaProducer) PublishSimpleValueByte(Topic string, key string, value []byte) error {
	if err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &Topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
		Key:   []byte(key),
	}, nil); err != nil {
		return err
	}

	return nil
}

func (k *KafkaProducer) Publish(ctx context.Context, spanName, msgType string, key string, entity interface{}) error {
	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(entity); err != nil {
		return err
	}

	var kafkaKey struct {
		Id string `json:"id"`
	}
	kafkaKey.Id = key
	bKey, err := json.Marshal(&kafkaKey)
	if err != nil {
		return err
	}
	topicUserFullInfo := "user_full_info"
	if err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topicUserFullInfo,
			Partition: kafka.PartitionAny,
		},
		Value: b.Bytes(),
		Key:   bKey,
	}, nil); err != nil {
		return err
	}

	return nil
}
