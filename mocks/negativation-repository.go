package mocks

import (
	di "github.com/klasrak/data-integration"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type MockNegativationRepository struct {
	mock.Mock
}

func (nr *MockNegativationRepository) InsertOne(n di.Negativation) error {
	ret := nr.Called(n)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (nr *MockNegativationRepository) InsertMany(nList []di.Negativation) error {
	ret := nr.Called(nList)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (nr *MockNegativationRepository) Update(id string, n *bson.M) (di.Negativation, error) {
	ret := nr.Called(id, n)

	var r0 di.Negativation

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(di.Negativation)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (nr *MockNegativationRepository) Delete(id string) error {
	ret := nr.Called(id)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (nr *MockNegativationRepository) GetOne(customerDocument string) ([]di.Negativation, error) {
	ret := nr.Called(customerDocument)

	var r0 []di.Negativation

	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]di.Negativation)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (nr *MockNegativationRepository) GetByID(id string) (di.Negativation, error) {
	ret := nr.Called(id)

	var r0 di.Negativation

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(di.Negativation)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (nr *MockNegativationRepository) GetAll() ([]di.Negativation, error) {
	ret := nr.Called()

	var r0 []di.Negativation

	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]di.Negativation)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
