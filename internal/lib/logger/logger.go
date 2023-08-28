package logger

import (
	"log"
)

func ErrorWrap(op, msg string) {
	log.Printf("%s: %s", op, msg)
}
