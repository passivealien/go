package models

type RetrieveResponse struct {
	Code    int32 `json:"code"`
	Status  string
	Message string `json:"message,omitempty"`
	Name    string `json:"name"`
	//Data    interface{} `json:",omitempty"`
}
