package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
	"github.com/klasrak/data-integration/api/helpers"
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

// Paths Information

// Fetch godoc
// @Summary Fetch data from Legacy API
// @Description Fetch data from Legacy API and saves into mongodb
// @ID Fetch
// @Consume application/json
// @Produce json
// @Success 200 {object} []helpers.Negativation
// @Failure 400 {object} helpers.HTTPError
// @Router /negativations/fetch [get]
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
		helpers.NewError(c, http.StatusBadRequest, errors.New("legacy API data already fetched and there is no new data to save"))
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetAll godoc
// @Summary Get all negativations
// @Description Get all negativations from database
// @ID GetAll
// @Produce  json
// @Success 200 {object} []di.Negativation
// @Failure 404 {object} helpers.HTTPError
// @Router /negativations/get [get]
func (n *negativationHandler) GetAll(c *gin.Context) {
	result, err := n.repo.GetAll()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.NewError(c, http.StatusNotFound, errors.New("no negativations registered so far"))
			return
		} else {
			helpers.NewError(c, http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get godoc
// @Summary Get negativations
// @Description Get all negativations from a documentNumber
// @ID Get
// @Produce json
// @Param customerDocument path string true "Customer document (CPF)"
// @Success 200 {object} []di.Negativation
// @Failure 400 {object} helpers.HTTPError
// @Failure 404 {object} helpers.HTTPError
// @Failure 500 {object} helpers.HTTPError
// @Router /negativations/get{customerDocument} [get]
func (n *negativationHandler) Get(c *gin.Context) {
	customerDocument := c.Param("customerDocument")

	if customerDocument == "" {
		helpers.NewError(c, http.StatusBadRequest, errors.New("invalid document number"))
		return
	}

	result, err := n.repo.GetOne(customerDocument)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.NewError(c, http.StatusNotFound, errors.New("negativation not found"))
			return
		} else {
			helpers.NewError(c, http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, result)
}

// GetByID godoc
// @Summary Get a negativation by ID
// @Description Get a negativation by ID
// @ID GetByID
// @Produce json
// @Param id path string true "Negativation ID"
// @Success 200 {object} di.Negativation
// @Failure 404 {object} helpers.HTTPError
// @Failure 500 {object} helpers.HTTPError
// @Router /negativations/get/{id} [get]
func (n *negativationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := n.repo.GetByID(id)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.NewError(c, http.StatusNotFound, errors.New("negativation not found"))
			return
		} else {
			helpers.NewError(c, http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create godoc
// @Summary Create negativation
// @Description Create negativation
// @ID Create
// @Accept json
// @Produce json
// @Param negativation body helpers.Negativation true "Add negativation"
// @Success 200 {object} di.Negativation
// @Failure 400 {object} helpers.HTTPError
// @Failure 500 {object} helpers.HTTPError
// @Router /negativations/create [post]
func (n *negativationHandler) Create(c *gin.Context) {
	var negativation = di.Negativation{}

	err := c.Bind(&negativation)

	if err != nil {
		panic(err.Error())
	}

	err = n.repo.InsertOne(negativation)

	if mongo.IsDuplicateKeyError(err) {
		helpers.NewError(c, http.StatusBadRequest, errors.New("negativation already exists"))
		return
	}

	c.JSON(http.StatusCreated, negativation)
}

// Update godoc
// @Summary Update negativation
// @Description Update negativation
// @ID Update
// @Accept json
// @Produce json
// @Param id path int true "Negativation ID"
// @Param data body helpers.Negativation true "Data to update"
// @Success 200 {object} di.Negativation
// @Failure 400 {object} helpers.HTTPError
// @Failure 500 {object} helpers.HTTPError
// @Router /negativations/update/{id} [put]
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

	result, err := n.repo.Update(id, data)

	if err != nil {
		if err.Error() == "invalid id" {
			helpers.NewError(c, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		helpers.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete godoc
// @Summary Delete negativation
// @Description Delete negativation
// @ID Delete
// @Accept json
// @Produce json
// @Param id path int true "Negativation ID"
// @Success 200
// @Failure 500 {object} helpers.HTTPError
// @Router /negativations/delete/{id} [delete]
func (n *negativationHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := n.repo.Delete(id)

	if err != nil {
		helpers.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
