package main

import (
	"medsearch/config"
	"medsearch/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	config.UpdateDB(db)
	// Initialize Gin router
	r := gin.Default()
	user := handler.NewUserHandler()

	// Define routes
	r.GET("/users/:id", user.Get)
	r.POST("/users", user.Create)
	// r.PUT("/users/:id", updateUserHandler)
	// r.DELETE("/users/:id", deleteUserHandler)

	// Start the server
	r.Run(":3007")
}
