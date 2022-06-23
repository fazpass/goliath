package queue

import "context"

type Config struct {
	Driver string
	Host   string
}

func Init(ctx context.Context, config Config) interface{} {
	switch config.Driver {
	case "kafka":
		return InitKafka([]string{config.Host})
	default:
		return nil
	}
}
