package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	Fetch(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
