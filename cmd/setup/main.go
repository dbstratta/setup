package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/strattadb/setup/internal/app/ubuntu"
)

func main() {
	operatingSystem := detectOS()

	switch operatingSystem {
	case "darwin":
		log.Fatal("macOS not implemented")
	case "linux":
		ubuntu.Setup()
	}
}

func detectOS() string {
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

		os.Exit(1)
	}

	return operatingSystem
}
