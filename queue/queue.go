package queue

import (
	"context"
	"errors"
)

type Config struct {
	Driver  string
	Host    string
	Kind    string
	GroupId string
}

func Init(ctx context.Context, config Config) (interface{}, error) {
	switch config.Driver {
	case "kafka":
		switch config.Kind {
		case "producer":
			return InitKafkaProducer([]string{config.Host})
		case "consumer":
			return InitKafkaConsumer([]string{config.Host})
		case "consumer_group":
			return InitKafkaConsumerGroup([]string{config.Host}, config.GroupId)
		default:
			return nil, errors.New("queue kind does not exists")
		}
	default:
		return nil, errors.New("queue driver does not exists")
	}
}
