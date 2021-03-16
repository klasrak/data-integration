package cryptography

import (
	"testing"

	"github.com/klasrak/data-integration/internal/converter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var stb = converter.StringToBytes

type encrypterMock struct {
	mock.Mock
}

func (e *encrypterMock) Encrypt(value string) []byte {
	args := e.Called(value)
	return args.Get(0).([]byte)
}

func (e *encrypterMock) Decrypt(data []byte) string {
	args := e.Called(data)
	return args.String(0)
}

func TestEncryptSuccess(t *testing.T) {
	sut := new(encrypterMock)

	// Set expectations
	sut.On("Encrypt", mock.Anything).Return(stb("any_value"), nil)

	assert.NotPanics(t, func() {
		encrypted := sut.Encrypt("any_value")
		assert.Equal(t, stb("any_value"), encrypted, "Values should be equal")
	}, "should not panic")

	sut.AssertExpectations(t)
}

func TestEncryptFail(t *testing.T) {
	sut := new(encrypterMock)

	sut.On("Encrypt", mock.Anything).Return(nil)

	assert.Panics(t, func() {
		sut.Encrypt("any_value")
	}, "should panic")

	sut.AssertExpectations(t)
}
