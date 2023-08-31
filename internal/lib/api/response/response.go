package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

var (
	DuplicateErr    = errors.New("already exists")
	NotFoundErr     = errors.New("not found")
	InvalidJson     = errors.New("invalid json")
	NotAddedSlugs   = errors.New("added slugs does not exist")
	NotDeletedSlugs = errors.New("deleted slugs does not exist")
)

var clientsErrs []error = []error{
	DuplicateErr, NotFoundErr, NotAddedSlugs, NotDeletedSlugs,
}

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

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

func ValidationError(w http.ResponseWriter, err validator.ValidationErrors) {
	var errorMsgs []string
	for _, err := range err {
		switch err.Tag() {
		case "required":
			errorMsgs = append(errorMsgs, fmt.Sprintf("field %s is required", err.Field()))
		case "notblank":
			errorMsgs = append(errorMsgs, fmt.Sprintf("field %s cannot be blank", err.Field()))
		}
	}
	w.Write(Error(strings.Join(errorMsgs, ",")))
}

func isClientError(e error) bool {
	for _, err := range clientsErrs {
		if err == e {
			return true
		}
	}
	return false
}

func SendError(w http.ResponseWriter, err error) {
	if isClientError(err) {
		w.Write(Error(err.Error()))
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
