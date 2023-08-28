package helpers

import (
	"encoding/json"
	"io"
)

func ParseJson(body io.ReadCloser, data interface{}) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
