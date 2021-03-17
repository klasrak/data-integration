package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
	"github.com/klasrak/data-integration/api/helpers"
	"github.com/klasrak/data-integration/jwt"
	rep "github.com/klasrak/data-integration/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

var tokenManager = jwt.TokenManager{}

type authHandler struct {
	r         rep.UsersRepository
	jwtSecret string
}

func NewAuthHandler(r rep.UsersRepository, jwtSecret string) *authHandler {
	return &authHandler{
		r:         r,
		jwtSecret: jwtSecret,
	}
}

// Paths Information

// Login godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Login
// @Accept  json
// @Produce  json
// @Param credentials body helpers.Login true "User credentials"
// @Success 200 {object} helpers.Tokens
// @Failure 404 {object} helpers.HTTPError
// @Router /auth/login [post]
func (auth *authHandler) Login(c *gin.Context) {
	var u di.User

	if err := c.Bind(&u); err != nil {
		helpers.NewError(c, http.StatusUnprocessableEntity, err)
		return
	}

	user, err := auth.r.FindByEmail(u.Email)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.NewError(c, http.StatusNotFound, errors.New("user not found"))
			return
		} else {
			helpers.NewError(c, http.StatusInternalServerError, err)
			return
		}
	}

	if user.Email != u.Email || user.Password != u.Password {
		helpers.NewError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	ts, err := tokenManager.CreateToken(user.ID, user.Email, auth.jwtSecret)

	if err != nil {
		helpers.NewError(c, http.StatusUnprocessableEntity, err)
		return
	}

	tokens := helpers.Tokens{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

func (auth *authHandler) Logout(c *gin.Context) {
	// TODO: delete token from redis (or another) store
	c.JSON(http.StatusOK, "successfully logged out")
}

func (auth *authHandler) Refresh(c *gin.Context) {
	// TODO: create refresh token logic
	tokens := map[string]string{
		"accessToken":  "validAccessToken",
		"refreshToken": "validRefreshToken",
	}

	c.JSON(http.StatusCreated, tokens)
}
