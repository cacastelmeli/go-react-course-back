package infrastructure_shared

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient = setupMongoClient()
var MongoDatabase = MongoClient.Database("test")

func setupMongoClient() *mongo.Client {
	uri := "mongodb://127.0.0.1:27017/test"

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return client
}
