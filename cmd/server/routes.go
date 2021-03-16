package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/api/handlers"
	"github.com/klasrak/data-integration/api/middlewares"
	"github.com/klasrak/data-integration/api/routes"
	rp "github.com/klasrak/data-integration/repositories"
)

func (s *Server) InitRoutes() {

	jwtMiddleware := middlewares.TokenAuthMiddleware(s.Env.JWT_SECRET)

	usersRepository := rp.NewUsersRepository(s.MongoClient)
	negativationRepository := rp.NewNegativationRepository(s.MongoClient)

	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World",
		})
	})

	api := s.Router.Group("/api")
	{
		v1 := api.Group("v1/")
		{
			routes.AuthRoutes(v1, nil, handlers.NewAuthHandler(usersRepository, s.Env.JWT_SECRET))
			routes.UsersRouter(v1, jwtMiddleware, handlers.NewUsersHandler(usersRepository))
			routes.NegativationRoutes(v1, jwtMiddleware, handlers.NewNegativationHandler(negativationRepository))
		}
	}
}
