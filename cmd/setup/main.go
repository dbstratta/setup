package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/strattadb/setup/internal/app/ubuntu"
	"github.com/strattadb/setup/internal/pkg/helpers"
)

func main() {
	operatingSystem, err := detectOS()
	helpers.ExitIfError(err)

	switch operatingSystem {
	case "darwin":
		log.Fatal("macOS not implemented")
	case "linux":
		ubuntu.Setup()
	}
}

func detectOS() (string, error) {
	const operatingSystem = runtime.GOOS

	switch operatingSystem {
	case "darwin":
		fmt.Println("os: macOS")
	case "linux":
		fmt.Println("os: Linux")
	default:
		fmt.Fprintf(
			os.Stderr,
			"setup: OS %s not supported\n",
			operatingSystem)

		return "", errors.New("OS not supported")
	}

	return operatingSystem, nil
}
