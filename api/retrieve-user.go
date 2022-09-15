package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RetrieveUser(c *gin.Context) {

	// TODO: 📌 Handle insert records in creating users

	c.JSON(http.StatusOK, gin.H{
		"TODO": "Retrieve user " + c.Param("user"),
	})
}
