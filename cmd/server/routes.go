package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/api/handlers"
	"github.com/klasrak/data-integration/api/routes"
	rp "github.com/klasrak/data-integration/repositories"
)

func (s *Server) InitRoutes() {

	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World",
		})
	})

	api := s.Router.Group("/api")
	{
		v1 := api.Group("v1/")
		{
			routes.UsersRouter(v1, handlers.NewUsersHandler(rp.NewUsersRepository(s.MongoClient)))
			routes.NegativationRoutes(v1, handlers.NewNegativationHandler(rp.NewNegativationRepository(s.MongoClient)))
		}
	}
}
