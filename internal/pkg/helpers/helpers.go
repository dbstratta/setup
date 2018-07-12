package helpers

import (
	"log"
	"os"
)

// LogAndExitIfError logs and exits the program if `err`
// is not `nil`.
func LogAndExitIfError(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

// ExitIfError exits the program if `err` is not `nil`.
func ExitIfError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
