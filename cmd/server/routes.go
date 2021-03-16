package server

import (
	"github.com/klasrak/data-integration/api/handlers"
	"github.com/klasrak/data-integration/api/middlewares"
	"github.com/klasrak/data-integration/api/routes"
	rp "github.com/klasrak/data-integration/repositories"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (s *Server) InitRoutes() {

	jwtMiddleware := middlewares.TokenAuthMiddleware(s.Env.JWT_SECRET)

	usersRepository := rp.NewUsersRepository(s.MongoClient)
	negativationRepository := rp.NewNegativationRepository(s.MongoClient)

	api := s.Router.Group("/api")
	{
		v1 := api.Group("v1/")
		{
			routes.AuthRoutes(v1, nil, handlers.NewAuthHandler(usersRepository, s.Env.JWT_SECRET))
			routes.UsersRouter(v1, jwtMiddleware, handlers.NewUsersHandler(usersRepository))
			routes.NegativationRoutes(v1, jwtMiddleware, handlers.NewNegativationHandler(negativationRepository))
		}
	}

	s.Router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
