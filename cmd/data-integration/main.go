package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/api/handlers"
	"github.com/klasrak/data-integration/config"
	"github.com/klasrak/data-integration/mongo"
	"github.com/klasrak/data-integration/repositories"
)

func main() {
	env := config.New()
	mongoClient, err := mongo.GetClient(env.MONGO_URI)

	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal("Failed to connect to mongodb")
	}

	negativationRepository := repositories.NewNegativationRepository(mongoClient)
	negativationHandler := handlers.NewNegativationHandler(&negativationRepository)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World",
		})
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/fetch", negativationHandler.Fetch)
		}
	}

	router.Run(":8080")

}
