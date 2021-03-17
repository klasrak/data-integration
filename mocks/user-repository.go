package mocks

import (
	di "github.com/klasrak/data-integration"
	"github.com/stretchr/testify/mock"
)

type MockUsersRepository struct {
	mock.Mock
}

func (r *MockUsersRepository) FindByEmail(email string) (di.User, error) {
	ret := r.Called(email)

	var r0 di.User

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(di.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (r *MockUsersRepository) FindAll() ([]di.User, error) {
	ret := r.Called()

	var r0 []di.User

	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]di.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
