package mongo

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func GetClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb:mongo:27017")

		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
