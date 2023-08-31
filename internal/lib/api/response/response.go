package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	DuplicateErr = errors.New("already exists")
	NotFoundErr  = errors.New("not found")
	InvalidJson  = errors.New("invalid json")
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

func SendError(w http.ResponseWriter, err error) {
	if errors.Is(err, DuplicateErr) {
		w.Write(Error(DuplicateErr.Error()))
	} else if errors.Is(err, NotFoundErr) {
		w.Write(Error(NotFoundErr.Error()))
	} else if errors.Is(err, InvalidJson) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(InvalidJson.Error()))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Send(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(OK(data))
}
