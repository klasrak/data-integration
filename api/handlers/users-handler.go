package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rep "github.com/klasrak/data-integration/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersHandler struct {
	repo rep.UsersRepository
}

func NewUsersHandler(r rep.UsersRepository) *usersHandler {
	return &usersHandler{
		repo: r,
	}
}

func (u *usersHandler) FindByEmail(c *gin.Context) {
	email := c.Param("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing param email",
		})
		return
	}

	user, err := u.repo.FindByEmail(email)

	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *usersHandler) Find(c *gin.Context) {
	result, err := u.repo.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
