package main

import (
	"github.com/klasrak/data-integration/cmd/server"

	_ "github.com/klasrak/data-integration/docs"
)

// @title Data Integration API
// @version 1.0
// @description This is a Rest API to consume from a Legacy API and serve.
// @contact.name Danilo Augusto
// @contact.url https://daniloaugusto.dev
// @contact.email dasfcm@gmail.com
// @license.name MIT
// @host localhost:8080
// @BasePath /api/v1
func main() {

	// Initialize API
	server.Run()

}
