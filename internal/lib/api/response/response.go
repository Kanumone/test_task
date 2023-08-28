package response

import (
	"encoding/json"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func OK(data interface{}) []byte {
	res, _ := json.Marshal(Response{
		Status: StatusOk,
		Data:   data,
	})
	return res
}

func Error(msg string) []byte {
	res, _ := json.Marshal(&Response{
		Status:  StatusError,
		Message: msg,
	})
	return res
}
