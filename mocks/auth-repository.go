package mocks

import (
	di "github.com/klasrak/data-integration"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (a *MockAuthRepository) CreateAuth(userID string, td *di.TokenDetails) error {
	ret := a.Called(userID, td)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (a *MockAuthRepository) FetchAuth(tokenUUID string) (string, error) {
	ret := a.Called(tokenUUID)

	var r0 string

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (a *MockAuthRepository) DeleteTokens(authD *di.AccessDetails) error {
	ret := a.Called(authD)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (a *MockAuthRepository) DeleteRefresh(refreshUUID string) error {
	ret := a.Called(refreshUUID)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
