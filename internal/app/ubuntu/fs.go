package ubuntu

import (
	"os"
	"os/user"
	"strings"

	"github.com/strattadb/setup/internal/pkg/helpers"
)

var folderPaths = []string{
	"~/projects",
	"~/work",
	"~/cs",
}

func initializeFileSystem() {
	currentUser, err := user.Current()
	helpers.ExitIfError(err)

	homeDir := currentUser.HomeDir

	for _, folderPath := range folderPaths {
		path := strings.Replace(folderPath, "~", homeDir, -1)
		os.MkdirAll(path, os.ModePerm)
	}
}
