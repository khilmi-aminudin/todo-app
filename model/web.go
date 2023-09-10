package model

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type HttpErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const ErrConstraint = "1451"

type EmptyStruct struct{}
