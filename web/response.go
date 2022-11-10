package web

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `jsjon:"message"`
	Data    interface{} `json:"data"`
}
