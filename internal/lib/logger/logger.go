package logger

import "fmt"

func ErrorWrap(op, msg string) error {
	return fmt.Errorf("%s: %s", op, msg)
}
