package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	// TODO: 📌 Handle delete users

	c.JSON(http.StatusOK, gin.H{
		"TODO": "Delete user",
	})
}
