package models

type RetrieveResponse struct {
	Code      int32 `json:"code"`
	Status    string
	Message   string `json:"message,omitempty"`
	ID        string `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	BirthDate string `json:"BirthDate"`
	Address   string `json:"Address"`
	JobTitle  string `json:"JobTitle"`
	EMail     string `json:"EMail"`
	Password  string `json:"Password"`
	//Data    interface{} `json:",omitempty"`

}
