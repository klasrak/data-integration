package repositories

import (
	"context"
	"time"

	di "github.com/klasrak/data-integration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NegativationRepository interface {
	InsertOne(n di.Negativation) error
	InsertMany(n []di.Negativation) error
	Update(n, newN di.Negativation) error
	Delete(customerDocument string) error

	GetOne(customerDocument string) (di.Negativation, error)
	GetAll() ([]di.Negativation, error)
}

type negativationRepository struct {
	collection mongo.Collection
}

func NewNegativationRepository(client *mongo.Client) NegativationRepository {
	db := client.Database("di_db")
	collection := db.Collection("negativations")

	mod := mongo.IndexModel{
		Keys: bson.M{
			"contract": 1,
		}, Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection.Indexes().CreateOne(ctx, mod)

	return &negativationRepository{
		collection: *collection,
	}
}

func (nr *negativationRepository) InsertOne(n di.Negativation) error {
	_, err := nr.collection.InsertOne(context.TODO(), n)

	if err != nil {
		return err
	}

	return nil
}

func (nr *negativationRepository) InsertMany(nList []di.Negativation) error {
	insertableList := make([]interface{}, len(nList))

	for i, v := range nList {
		r, _ := nr.GetOne(v.CustomerDocument)
		if r.ID == "" {
			insertableList[i] = v
		}
	}

	if insertableList[0] == nil {
		return nil
	}

	_, err := nr.collection.InsertMany(context.TODO(), insertableList)

	if err != nil {
		return err
	}

	return nil
}

func (nr *negativationRepository) Update(n, newN di.Negativation) error {
	nByte, err := bson.Marshal(newN)

	if err != nil {
		return err
	}

	var update bson.M
	err = bson.Unmarshal(nByte, &update)

	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "customerDocument", Value: n.CustomerDocument}}

	result := nr.collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})

	if result.Err() == mongo.ErrNoDocuments {
		return result.Err()
	}

	return nil
}

func (nr *negativationRepository) Delete(customerDocument string) error {
	filter := bson.D{primitive.E{Key: "customerDocument", Value: customerDocument}}

	_, err := nr.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (nr *negativationRepository) GetOne(customerDocument string) (di.Negativation, error) {
	result := di.Negativation{}

	filter := bson.D{primitive.E{Key: "customerDocument", Value: customerDocument}}

	err := nr.collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (nr *negativationRepository) GetAll() ([]di.Negativation, error) {
	filter := bson.D{{}}
	results := []di.Negativation{}

	cur, findError := nr.collection.Find(context.TODO(), filter)

	if findError != nil {
		return results, findError
	}

	for cur.Next(context.TODO()) {
		t := di.Negativation{}
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
