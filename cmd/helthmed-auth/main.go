package main

import (
	"helthmed-auth/internal/auth/infrastructure/db"
	"helthmed-auth/internal/auth/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Set up Gin router
	router := gin.Default()

	// Initialize HTTP handlers and routes
	http.InitRoutes(router)

	// Run the server
	router.Run(":8080")
}
