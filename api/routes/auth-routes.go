package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/api/handlers"
)

func AuthRoutes(router *gin.RouterGroup, middleware gin.HandlerFunc, h handlers.AuthHandler) {
	group := router.Group("/auth")
	if middleware != nil {
		group.Use(middleware)
	}
	{
		group.POST("/login", h.Login)
		group.POST("/logout", h.Logout)
		group.POST("/refresh-token", h.Refresh)
	}
}
