package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(ctx context.Context, source string) *mongo.Client {

	var client, err = mongo.Connect(ctx, options.Client().ApplyURI(source))
	if err != nil {
		panic(err)
	}

	return client

}
