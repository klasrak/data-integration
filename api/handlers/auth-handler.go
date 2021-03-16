package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	di "github.com/klasrak/data-integration"
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
// @Consume application/json
// @Produce json
// @Param body di.User "Login"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.Response
// @Router /auth/login [post]
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

	if user.Email != u.Email || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	ts, err := tokenManager.CreateToken(user.ID, user.Email, auth.jwtSecret)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokens := map[string]string{
		"accessToken":  ts.AccessToken,
		"refreshToken": ts.RefreshToken,
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
