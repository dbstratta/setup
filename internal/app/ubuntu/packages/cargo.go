package packages

import (
	"os/exec"

	"github.com/strattadb/setup/internal/pkg/helpers"
)

var cargoPackages = [50]string{
	"exa",
	"ripgrep",
	"bat",
	"fd-find",
}

func installCargoPackages() {
	for _, cargoPackage := range cargoPackages {
		err := exec.Command("cargo", "install", cargoPackage).Run()
		helpers.LogAndExitIfError(err)
	}
}
