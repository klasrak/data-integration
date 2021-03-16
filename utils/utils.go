package utils

import "go.mongodb.org/mongo-driver/bson"

func ToDoc(v interface{}) (doc *bson.M, err error) {
	data, err := bson.Marshal(v)

	if err != nil {
		panic(err.Error())
	}

	err = bson.Unmarshal(data, &doc)
	return
}
