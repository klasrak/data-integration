package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
	"github.com/klasrak/data-integration/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	httpClientMock := mocks.MockClient{}

	successJsonRsponse := `[
		{
		   "companyDocument":"59291534000167",
		   "companyName":"ABC S.A.",
		   "customerDocument":"51537476467",
		   "value":1235.23,
		   "contract":"bc063153-fb9e-4334-9a6c-0d069a42065b",
		   "debtDate":"2015-11-13T20:32:51-03:00",
		   "inclusionDate":"2020-11-13T20:32:51-03:00"
		},
		{
		   "companyDocument":"77723018000146",
		   "companyName":"123 S.A.",
		   "customerDocument":"51537476467",
		   "value":400,
		   "contract":"5f206825-3cfe-412f-8302-cc1b24a179b0",
		   "debtDate":"2015-10-12T20:32:51-03:00",
		   "inclusionDate":"2020-10-12T20:32:51-03:00"
		}
	 ]`

	t.Run("Success", func(t *testing.T) {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(successJsonRsponse)))
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		}
		req, _ := http.NewRequest(http.MethodGet, "http://legacy-api:3333/negativations", nil)

		res, err := httpClientMock.Do(req)

		body, _ := ioutil.ReadAll(res.Body)
		var result []di.Negativation
		json.Unmarshal(body, &result)

		assert.NoError(t, err, "No error")
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.StatusCode)
		assert.JSONEq(t, successJsonRsponse, string(body))

		mockNegativationRepository := new(mocks.MockNegativationRepository)
		mockNegativationRepository.On("InsertMany", result).Return(nil)

		err = mockNegativationRepository.InsertMany(result)
		assert.NoError(t, err, "no error")
		mockNegativationRepository.AssertExpectations(t)
	})

	t.Run("Error in legacy API", func(t *testing.T) {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return nil, errors.New("someError")
		}

		req, err := http.NewRequest(http.MethodGet, "http://legacy-api:3333/negativations", nil)

		assert.NotNil(t, req)
		assert.NoError(t, err)

		res, err := httpClientMock.Do(req)

		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, "someError", err.Error())

	})

	t.Run("Error in repository", func(t *testing.T) {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(successJsonRsponse)))
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		}
		req, _ := http.NewRequest(http.MethodGet, "http://legacy-api:3333/negativations", nil)

		res, _ := httpClientMock.Do(req)

		body, _ := ioutil.ReadAll(res.Body)
		var result []di.Negativation
		json.Unmarshal(body, &result)

		mockNegativationRepository := new(mocks.MockNegativationRepository)
		mockNegativationRepository.On("InsertMany", result).Return(errors.New("someError"))

		err := mockNegativationRepository.InsertMany(result)

		assert.Error(t, err)
		assert.Equal(t, "someError", err.Error())
		mockNegativationRepository.AssertExpectations(t)
	})
}
