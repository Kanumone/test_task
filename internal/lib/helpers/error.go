package helpers

import (
	"fmt"
	"log"
)

const (
	DuplicateSlug = "DUPLICATE SLUG"
)

func Wrap(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}

func LogErr(prefix string, err error) {
	log.Printf("%s: %s", prefix, err.Error())
}
