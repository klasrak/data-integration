package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/jwt"
)

func TokenAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.TokenValid(c.Request, jwtSecret)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
