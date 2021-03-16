package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
	rep "github.com/klasrak/data-integration/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type authHandler struct {
	r rep.UsersRepository
}

func NewAuthHandler(r rep.UsersRepository) *authHandler {
	return &authHandler{
		r: r,
	}
}

func (auth *authHandler) Login(c *gin.Context) {
	var u di.User

	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid body provided",
		})
		return
	}

	user, err := auth.r.FindByEmail(u.Email)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	if u.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	tokens := map[string]string{
		"accessToken":  "validAccessToken",
		"refreshToken": "validRefreshToken",
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tokens,
	})

}
