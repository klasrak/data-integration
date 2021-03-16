package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetClient(uri string) (*mongo.Client, error) {
	// Database Config
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		panic(err.Error())
	}

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	firstUser := bson.M{"email": "admin@mail.com", "password": "abcd1234"}

	_, err = client.Database("di_db").Collection("users").InsertOne(context.TODO(), firstUser)

	if err != nil {
		panic(err.Error())
	}

	return client, err
}
