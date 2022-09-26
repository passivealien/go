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

func DeleteUser(c *gin.Context) {

	var deletePayload models.DeleteInput
	var response models.RetrieveResponse

	err := c.ShouldBindJSON(&deletePayload)

	if err == nil {
		response.Status = "Successfull Delete"
		response.Code = 200
		response.Message = "Deleted User ID number: " + deletePayload.UserIDToDelete

	} else {
		response.Status = "Failed to Delete"
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
	UserIDToDelete := deletePayload.UserIDToDelete
	deleteID, _ := strconv.Atoi(UserIDToDelete)

	sqlStatement := `
	DELETE FROM "Users"
	WHERE "User ID" = $1;`
	_, err = db.Exec(sqlStatement, deleteID)
	if err != nil {
		panic(err)
	}

}

func CheckErrorDelete(err error) {
	if err != nil {
		panic(err)
	}
}
