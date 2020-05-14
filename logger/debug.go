package logger

import (
	"log"
	"os"
)

var debug = os.Getenv("LDAP_DEBUG") != ""

func Debug(format string, args ...interface{}) {
	if debug {
		log.Printf(format, args...)
	}
}
