package cmd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go_di_template/pkg/database"
)

func InitMongoDB(ctx context.Context, mongo *database.MongoDB) (*mongo.Database, *mongo.Client) {
	client, err := mongo.Connect()
	if err != nil {
		message := fmt.Sprintf("Cannot connect to mongoDB: %v", err)
		panic(message)
	}
	if err = client.Ping(ctx, nil); err != nil {
		message := fmt.Sprintf("Cannot connect to mongoDB: %v", err)
		panic(message)
	}
	return client.Database(mongo.Database), client
}
