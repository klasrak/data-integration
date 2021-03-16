package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
	rep "github.com/klasrak/data-integration/repositories"
	"github.com/klasrak/data-integration/utils"
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
	legacyApiUrl := "http://legacy-api:3333/negativations"

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

	if mongo.IsDuplicateKeyError(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Legacy API data already fetched",
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

	result, err := n.repo.GetOne(customerDocument)

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

func (n *negativationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := n.repo.GetByID(id)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "negativation not found",
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

func (n *negativationHandler) Create(c *gin.Context) {
	var negativation = di.Negativation{}

	err := c.Bind(&negativation)

	if err != nil {
		panic(err.Error())
	}

	err = n.repo.InsertOne(negativation)

	if mongo.IsDuplicateKeyError(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "negativation already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": negativation,
	})
}

func (n *negativationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var update = make(map[string]interface{})

	err := c.Bind(&update)

	if err != nil {
		panic(err.Error())
	}

	data, err := utils.ToDoc(update)

	if err != nil {
		panic(err.Error())
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid customer document",
		})
		return
	}

	err = n.repo.Update(id, data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": update,
	})
}

func (n *negativationHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := n.repo.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
