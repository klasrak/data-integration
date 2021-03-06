package repositories

import (
	"context"
	"time"

	di "github.com/klasrak/data-integration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersRepository interface {
	FindByEmail(email string) (di.User, error)
	FindAll() ([]di.User, error)
}

type usersRepository struct {
	collection mongo.Collection
}

func NewUsersRepository(client *mongo.Client) UsersRepository {
	db := client.Database("di_db")
	collection := db.Collection("users")

	mod := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		}, Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection.Indexes().CreateOne(ctx, mod)

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

func (u *usersRepository) FindAll() ([]di.User, error) {
	filter := bson.D{{}}
	results := []di.User{}

	cur, err := u.collection.Find(context.TODO(), filter)

	if err != nil {
		return results, err
	}

	for cur.Next(context.TODO()) {
		t := di.User{}
		err := cur.Decode(&t)

		if err != nil {
			return results, err
		}

		results = append(results, t)
	}

	cur.Close(context.TODO())

	if len(results) == 0 {
		return results, mongo.ErrNoDocuments
	}

	return results, nil
}
