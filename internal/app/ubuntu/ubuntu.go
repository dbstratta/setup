package ubuntu

import "github.com/strattadb/setup/internal/app/ubuntu/packages"

// Setup initializes an Ubuntu machine.
func Setup() {
	packages.InstallPackages()
	initializeFileSystem()
	copyDotfiles()
}
