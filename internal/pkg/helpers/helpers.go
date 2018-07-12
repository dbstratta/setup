package helpers

import (
	"log"
	"os"
	"os/exec"
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

// RunBashCommand runs a command with `bash -c ...`.
func RunBashCommand(cmdStr string) error {
	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()

	return err
}

// RunBashCommandAndLogAndExitIfError runs a command with
// `bash -c ...` and logs and exits if there's an error.
func RunBashCommandAndLogAndExitIfError(cmdStr string) {
	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()

	LogAndExitIfError(err)
}
