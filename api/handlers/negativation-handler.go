package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
	rep "github.com/klasrak/data-integration/repositories"
)

type negativationHandler struct {
	repository *rep.NegativationRepository
}

func NewNegativationHandler(r *rep.NegativationRepository) *negativationHandler {
	return &negativationHandler{
		repository: r,
	}
}

func (n *negativationHandler) Fetch(c *gin.Context) {
	client := http.Client{}
	legacyApiUrl := "http://localhost:3333/negativations"

	request, err := http.NewRequest(http.MethodGet, legacyApiUrl, nil)

	if err != nil {
		panic(err.Error())
	}

	res, err := client.Do(request)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var results []di.Negativation

	json.Unmarshal(body, &results)

	fmt.Println(results)

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}
