package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

func main() {
	_, err := detectOS()

	if err != nil {
		os.Exit(1)
	}
}

func detectOS() (string, error) {
	const operatingSystem = runtime.GOOS

	switch operatingSystem {
	case "darwin":
		fmt.Println("MacOS detected")
	case "linux":
		fmt.Println("Linux detected")
	default:
		fmt.Fprintf(os.Stderr, "setup: OS %s not supported\n", operatingSystem)
		return "", errors.New("OS not supported")
	}

	return operatingSystem, nil
}

func setupLinux() error {
	return nil
}
