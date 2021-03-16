package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/api/handlers"
)

func UsersRouter(router *gin.RouterGroup, middleware gin.HandlerFunc, h handlers.UserHandler) {
	group := router.Group("/users")
	group.Use(middleware)
	{
		group.GET("/get", h.Find)
		group.GET("/get/:email", h.FindByEmail)
	}
}
