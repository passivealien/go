package api

import (
	"clockedin/config"
	"clockedin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	var response models.GenericResponse

	response.Code = http.StatusOK
	response.Status = "success"
	response.Message = "running in " + config.GetEnvConfig("ENVIRONMENT") + " server..."

	c.JSON(http.StatusOK, response)
}
