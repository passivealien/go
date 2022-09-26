package models

type Input struct {
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	BirthDate string `json:"Birth Date"`
	Address   string `json:"Address"`
	JobTitle  string `json:"Job Title"`
	EMail     string `json:"E-Mail"`
	Password  string `json:"Password"`
	//UserID    int    `json:"User ID"`
}
