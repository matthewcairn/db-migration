package repo

import (
	"time"
	"transfer/conf"
)

const (
	generalQueryTimeout = 60 * time.Second
)

func newKafkaConfig() (host, topic string, err error) {
	host = conf.GetConfig().KafkaHost
	if host == "" {
		return "kafka", "", nil
	}
	return host, topic, nil
}
