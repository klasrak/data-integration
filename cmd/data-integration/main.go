package main

import (
	"github.com/klasrak/data-integration/cmd/server"

	_ "github.com/klasrak/data-integration/docs"
)

// @title Data Integration API
// @version 1.0
// @description API to fetch data from a legacy API, save on mongoDB and return data.

// @contact.name Danilo Augusto
// @contact.url https://daniloaugusto.dev
// @contact.email dasfcm@gmail.com

// @license.name MIT

// @BasePath /api/v1
func main() {

	// Initialize API
	server.Run()

}
