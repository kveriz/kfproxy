package client

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KFConfig struct {
	CFG *kafka.ConfigMap
}

type Config map[string]interface{}

func NewConfig(m map[string]interface{}) *KFConfig {
	cfg := &kafka.ConfigMap{}

	for k, v := range m {
		fmt.Println(k, v)
		cfg.SetKey(k, v)
	}

	return &KFConfig{
		CFG: cfg,
	}
}
