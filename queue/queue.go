package queue

import "context"

type Config struct {
	Driver   string
	Source   string
	Host     string
	Password string
}

func Init(ctx context.Context, config Config) interface{} {
	switch config.Driver {
	case "kafka":
		return InitKafka([]string{config.Host})
	default:
		return nil
	}
}
