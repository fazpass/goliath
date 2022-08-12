package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Driver       string
	Source       string
	Host         string
	Password     string
	Db           string
	MongoOptions options.ClientOptions
}

func Init(ctx context.Context, config Config) (interface{}, error) {
	switch config.Driver {
	case "postgres":
		return InitPostgresql(config.Driver, config.Source)
	case "mongodb":
		return InitMongoDB(ctx, &config.MongoOptions)
	case "redis":
		return InitRedis(ctx, config.Host, config.Password, config.Db)
	default:
		return nil, errors.New("database driver not found")
	}
}
