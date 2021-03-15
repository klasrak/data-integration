package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
	rep "github.com/klasrak/data-integration/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type negativationHandler struct {
	repo rep.NegativationRepository
}

func NewNegativationHandler(r rep.NegativationRepository) *negativationHandler {
	return &negativationHandler{
		repo: r,
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

	err = n.repo.InsertMany(results)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch legacy API data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func (n *negativationHandler) GetAll(c *gin.Context) {
	result, err := n.repo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (n *negativationHandler) Get(c *gin.Context) {
	customerDocument := c.Param("customerDocument")

	if customerDocument == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid customer document",
		})
		return
	}

	result, err := n.repo.GetOne(c.Param("customerDocument"))

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Negativation not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
