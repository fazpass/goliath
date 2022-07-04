package database

import (
	"context"
	"errors"
)

type Config struct {
	Driver   string
	Source   string
	Host     string
	Password string
	Db       string
}

func Init(ctx context.Context, config Config) (interface{}, error) {
	switch config.Driver {
	case "postgres":
		return InitPostgresql(config.Driver, config.Source)
	case "mongodb":
		return InitMongoDB(ctx, config.Source)
	case "redis":
		return InitRedis(ctx, config.Host, config.Password, config.Db)
	default:
		return nil, errors.New("driver not found")
	}
}
