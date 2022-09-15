package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	// TODO: 📌 Handle update user info

	c.JSON(http.StatusOK, gin.H{
		"TODO": "Update user",
	})
}
