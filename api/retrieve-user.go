package api

import (
	"clockedin/config"
	"clockedin/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RetrieveUser(c *gin.Context) {

	var retrievePayload models.RetrieveInput
	var response models.RetrieveResponse

	err := c.ShouldBindJSON(&retrievePayload)

	if err == nil {
		response.Status = "Successfull Retrieve"
		response.Code = 200
		response.Message = "Retrieve User ID number: " + retrievePayload.UserIDToRetrieve
		response.Name = "Name: "

	} else {
		response.Status = "Failed to Retrieve"
		response.Code = 500
		response.Message = err.Error()

	}

	c.JSON(http.StatusOK, response)

	if err != nil {
		panic(err)
	}

	host := config.GetEnvConfig("host")
	port := config.GetEnvConfig("port")
	user := config.GetEnvConfig("user")
	password := config.GetEnvConfig("password")
	dbname := config.GetEnvConfig("dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	UserIDToRetrieve := retrievePayload.UserIDToRetrieve
	retrieveID, _ := strconv.Atoi(UserIDToRetrieve)

	sqlStatement := `
	SELECT FROM "Users"
	WHERE "User ID" = $1;`
	_, err = db.Exec(sqlStatement, retrieveID)
	if err != nil {
		panic(err)
	}

}

func CheckErrorRetrieve(err error) {
	if err != nil {
		panic(err)
	}
}
