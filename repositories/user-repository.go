package repositories

import (
	"context"

	di "github.com/klasrak/data-integration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository interface {
	FindByEmail(email string) (di.User, error)
}

type usersRepository struct {
	collection mongo.Collection
}

func NewUsersRepository(client *mongo.Client) UsersRepository {
	db := client.Database("di_db")
	collection := db.Collection("users")

	return &usersRepository{
		collection: *collection,
	}
}

func (u *usersRepository) FindByEmail(email string) (di.User, error) {
	user := di.User{}

	filter := bson.M{"email": email}

	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}
