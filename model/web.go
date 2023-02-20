package model

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const ErrConstraint = "1451"

type EmptyStruct struct{}
