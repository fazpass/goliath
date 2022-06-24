package queue

import "context"

type Config struct {
	Driver string
	Host   string
	Kind   string
}

func Init(ctx context.Context, config Config) interface{} {
	switch config.Driver {
	case "kafka":
		switch config.Kind {
		case "producer":
			return InitKafkaProducer([]string{config.Host})
		case "consumer":
			return InitKafkaConsumer([]string{config.Host})
		default:
			return nil
		}
	default:
		return nil
	}
}
