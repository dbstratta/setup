package ubuntu

import (
	"os"
)

var folderPaths = []string{
	"~/projects",
	"~/work",
	"~/cs"}

func initializeFileSystem() {
	for _, folderPath := range folderPaths {
		os.MkdirAll(folderPath, os.ModePerm)
	}
}
