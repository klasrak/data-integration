package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/api/handlers"
)

func UsersRouter(router *gin.RouterGroup, h handlers.UserHandler) {
	group := router.Group("/users")
	{
		group.GET("/get")
		group.GET("/get/:email", h.FindByEmail)
	}
}
