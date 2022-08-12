package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(ctx context.Context, options *options.ClientOptions) (*mongo.Client, error) {

	var client, err = mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	return client, nil

}
