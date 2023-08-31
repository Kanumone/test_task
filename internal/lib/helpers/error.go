package helpers

import (
	"log"
)

func LogErr(prefix string, err error) {
	log.Printf("%s: %s", prefix, err.Error())
}
