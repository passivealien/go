package api

import (
	"clockedin/config"
	"clockedin/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RetrieveUser(c *gin.Context) {

	number := c.Param("number")
	intVar, _ := strconv.Atoi(number)

	var retrievePayload models.RetrieveInput
	var response models.RetrieveResponse

	err := c.ShouldBindJSON(&retrievePayload)

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

	sqlStatement := `
    SELECT "User ID","First Name", "Last Name", "Birth Date", "Address", "Job Title", "Email" FROM "Users"
    WHERE "User ID" = $1;`
	_, err = db.Exec(sqlStatement, intVar)
	if err != nil {
		panic(err)
	}

	var (
		id        int
		firstName string
		lastName  string
		birthDate string
		address   string
		jobTitle  string
		eMail     string
	)
	rows, err := db.Query(sqlStatement, intVar)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &birthDate, &address, &jobTitle, &eMail)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, firstName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		response.Status = "Successfull Retrieve"
		response.Code = 200
		response.Message = "Retrieve User ID number: "
		response.FirstName = firstName
		response.LastName = lastName
		response.BirthDate = birthDate
		response.Address = address
		response.JobTitle = jobTitle
		response.EMail = eMail

	} else {
		response.Status = "Failed to Retrieve"
		response.Code = 500
		response.Message = err.Error()

	}

	c.JSON(http.StatusOK, response)

	if err != nil {
		panic(err)
	}

}

func CheckErrorRetrieve(err error) {
	if err != nil {
		panic(err)
	}
}
