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

type UserHandler interface {
	FindByEmail(c *gin.Context)
	Find(c *gin.Context)
}

type AuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
}
