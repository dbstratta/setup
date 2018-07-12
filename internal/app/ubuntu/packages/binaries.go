package packages

import (
	"fmt"

	"github.com/strattadb/setup/internal/pkg/helpers"
)

func installBinaries() {
	installRust()
	installGo()
	installMinikube()
}

func installRust() {
	const rustupURL = "https://sh.rustup.rs"
	cmdStr := fmt.Sprintf("curl -sSf %s | sh", rustupURL)

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}

func installGo() {
	const goVersion = "1.10.3"
	goDownloadURL := fmt.Sprintf(
		"https://dl.google.com/go/go%s.linux-amd64.tar.gz", goVersion)

	cmdStr := fmt.Sprintf(` \
		curl %s && \
		tar -C /usr/local -xzf go%s.linux-amd64.tar.gz`,
		goDownloadURL, goVersion)

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}

func installMinikube() {
	const minikubeVersion = "0.28.0"
	minikubeDownloadURL := fmt.Sprintf(
		"https://storage.googleapis.com/minikube/releases/v%s/minikube-linux-amd64",
		minikubeVersion)

	cmdStr := fmt.Sprintf(` \
		curl -Lo minikube %s && \
		chmod +x minikube && \
		sudo mv minikube /usr/local/bin/`,
		minikubeDownloadURL)

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}
