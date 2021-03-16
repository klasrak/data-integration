package main

import (
	"github.com/klasrak/data-integration/cmd/server"
	"github.com/klasrak/data-integration/docs"
	// swagger embed files
	// gin-swagger middleware
)

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Data Integration Challenge - API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Initialize API
	server.Run()

}
