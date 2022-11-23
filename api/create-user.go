package api

import (
	// "clockedin/config"
	"clockedin/models"
	// "database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	// "fmt"
	// "math/rand"
)

func CreateUser(c *gin.Context) {

	var payload models.Input
	var response models.GenericResponse

	err := c.ShouldBindJSON(&payload)

	if err == nil {
		response.Status = "Successfull Input"
		response.Code = 200
		response.Message = "Hi, " + payload.FirstName

	} else {
		response.Status = "Failed"
		response.Code = 500
		response.Message = err.Error()

	}

	c.JSON(http.StatusOK, response)

	if err != nil {
		panic(err)
	}
	// 	host := config.GetEnvConfig("host")
	// 	port := config.GetEnvConfig("port")
	// 	user := config.GetEnvConfig("user")
	// 	password := config.GetEnvConfig("password")
	// 	dbname := config.GetEnvConfig("dbname")

	// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// 	db, err := sql.Open("postgres", psqlInfo)

	// 	FirstName := payload.FirstName
	// 	LastName := payload.LastName
	// 	BirthDate := payload.BirthDate
	// 	Address := payload.Address
	// 	JobTitle := payload.JobTitle
	// 	EMail := payload.EMail
	// 	Password := payload.Password

	// 	//UserID := rand.Int()
	// 	UserID := rand.Intn(9999)
	// 	fmt.Println(rand.Intn(200))
	// 	fmt.Println(rand.Intn(200))
	// 	fmt.Println(rand.Intn(200))
	// 	fmt.Println()

	// 	defer db.Close()

	// 	insertStmt := `insert into "Users"("FirstName", "LastName", "BirthDate", "Address", "JobTitle", "Email", "Password", "UserID") values($1, $2, $3, $4, $5, $6, $7, $8)`
	// 	_, e := db.Exec(insertStmt, FirstName, LastName, BirthDate, Address, JobTitle, EMail, Password, UserID)
	// 	CheckErrorCreate(e)
	// 	fmt.Println("Successfully connected!")

	// 	sqlStatement := `
	// 	DELETE FROM "Users"
	// 	WHERE "Password" = $1;`
	// 	_, err = db.Exec(sqlStatement, 1234)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// func CheckErrorCreate(err error) {
	// 	if err != nil {
	// 		panic(err)
	// 	}
}

//sample input
/*

{
	"First Name" : "Cardo",
	"Last Name" : "Dalisay",
	"Birth Date" : "January 1, 2000",
	"Address" : "Cebu City",
	"Job Title" : "Full Stack Developer Intern",
	"E-Mail" : "cards@jobtarget.com",
	"Password" : "1234",

}
*/
