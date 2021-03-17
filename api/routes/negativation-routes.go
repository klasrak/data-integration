package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/api/handlers"
)

func NegativationRoutes(router *gin.RouterGroup, middleware gin.HandlerFunc, h handlers.Handler) {
	group := router.Group("/negativations")
	group.Use(middleware)
	{
		group.GET("/fetch", h.Fetch)
		group.GET("/get", h.GetAll)
		group.GET("/get/:customerDocument", h.Get)
		group.GET("/get-id/:id", h.GetByID)
		group.POST("/create", h.Create)
		group.PUT("/update/:id", h.Update)
		group.DELETE("delete/:id", h.Delete)
	}
}
