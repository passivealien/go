package api

import (
	"net/http"

	"clockedin/config"
	"clockedin/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	//"fmt"

	_ "github.com/lib/pq"
)

func UpdateUser(c *gin.Context) {

	var payload models.Input
	var response models.GenericResponse

	err := c.ShouldBindJSON(&payload)
	if err == nil {
		response.Status = "success"
		response.Code = 200
		response.Message = "successfully updated!"

	} else {
		response.Status = "Failed"
		response.Code = 500
		response.Message = err.Error()
	}

	c.JSON(http.StatusOK, response)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", config.GetEnvConfig("host"), config.GetEnvConfig("port"), config.GetEnvConfig("user"), config.GetEnvConfig("password"), config.GetEnvConfig("dbname"))
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	updateStmt := `update "Information" set "BirthDate"=$1, "Password"=$2 where "ID"=$3`
	_, e := db.Exec(updateStmt, "March 10, 1995", "mariahl#", 4)

	if e != nil {
		panic(e)
	}
}
