package packages

// InstallPackages installs all packages.
func InstallPackages() {
	installAptPackages()
	installBinaries()
}
