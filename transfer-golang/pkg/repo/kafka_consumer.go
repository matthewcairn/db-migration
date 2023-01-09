package repo

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

func NewKafkaConsumer(groupID string) (*KafkaConsumer, error) {
	host, _, err := newKafkaConfig()
	if err != nil {
		return nil, fmt.Errorf("newKafkaConfig %w", err)
	}

	config := kafka.ConfigMap{
		"bootstrap.servers":  host,
		"group.id":           groupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	client, err := kafka.NewConsumer(&config)
	if err != nil {
		return nil, fmt.Errorf("kafka.NewConsumer %w", err)
	}

	if err := client.SubscribeTopics([]string{"mysql_server_name.subaid_01.Drug", "mysql_server_name.subaid_01.Patients"}, nil); err != nil {
		return nil, fmt.Errorf("client.Subscribe %w", err)
	}

	return &KafkaConsumer{
		Consumer: client,
	}, nil
}
