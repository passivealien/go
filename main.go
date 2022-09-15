package main

import (
	"clockedin/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// System
	router.GET("/api/health", api.HealthCheck)

	// User
	router.POST("/api/createUser", api.CreateUser)
	router.POST("/api/updateUser", api.UpdateUser)
	router.POST("/api/deleteUser", api.DeleteUser)

	// Lookup
	router.GET("/api/lookup/:user", api.RetrieveUser)

	router.Run(":80")
}
