package packages

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/strattadb/setup/internal/pkg/helpers"
)

var aptPackages = [50]string{
	"apt-transport-https",
	"ca-certificates",
	"curl",
	"software-properties-common",
	"build-essential",
	"cmake",
	"git",
	"vim",
	"fish",
	"nodejs",
	"python3.7",
	"docker-ce",
	"kubectl",
	"fonts-firacode",
	"xsel",
}

func installAptPackages() {
	removeAptRepositories()
	addAptRepositories()

	packages := strings.Join(aptPackages[:], " ")

	err := exec.Command("sudo", "apt-get", "update").Run()
	helpers.LogAndExitIfError(err)
	err = exec.Command("sudo", "apt-get", "-y", "install", packages).Run()
	helpers.LogAndExitIfError(err)

	postInstallDocker()
}

func removeAptRepositories() {

}

var aptRepositories = [1]string{"ppa:deadsnakes/ppa"}

func addAptRepositories() {
	repositories := strings.Join(aptRepositories[:], " ")

	cmd := exec.Command("sudo", "add-apt-repository", "-y", repositories)
	cmd.Run()

	addNodeJSAptRepository()
	addYarnAptRepository()
	addDockerAptRepository()
	addKubectlAptRepository()
}

func addNodeJSAptRepository() {
	const nodeVersion = "10"
	nodeSourceURL := fmt.Sprintf(
		"https://deb.nodesource.com/setup_%s.x",
		nodeVersion)

	cmdStr := fmt.Sprintf("curl -sL %s | sudo -E bash -", nodeSourceURL)

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}

func addYarnAptRepository() {
	const yarnPublicKeyURL = "https://dl.yarnpkg.com/debian/pubkey.gpg"
	const echoYarnList = `echo "deb https://dl.yarnpkg.com/debian/ stable main"`

	cmdStr := fmt.Sprintf(`
		curl -sS %s | sudo apt-key add - && \
		%s | sudo tee /etc/apt/sources.list.d/yarn.list`,
		yarnPublicKeyURL,
		echoYarnList)

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}

func addDockerAptRepository() {
	const dockerPublicKeyURL = "https://download.docker.com/linux/ubuntu/gpg"
	const echoDockerList = `echo "deb http://apt.kubernetes.io/ kubernetes-xenial main"`

	cmdStr := fmt.Sprintf(`
		curl -fsSL %s | sudo apt-key add - && \
		 %s | sudo tee /etc/apt/sources.list.d/kubernetes.list`,
		dockerPublicKeyURL,
		echoDockerList)

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}

func addKubectlAptRepository() {
	const kubectlPublicKeyURL = "https://packages.cloud.google.com/apt/doc/apt-key.gpg"
	const echoKubectlList = `echo "deb http://apt.kubernetes.io/ kubernetes-xenial main"`

	cmdStr := fmt.Sprintf(`
		curl -s %s | sudo apt-key add - && \
		 %s | sudo tee /etc/apt/sources.list.d/kubernetes.list`,
		kubectlPublicKeyURL,
		echoKubectlList)

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}

func postInstallDocker() {
	const cmdStr = `
		sudo groupadd docker && \
		sudo usermod -aG docker "${USER}"`

	helpers.RunBashCommandAndLogAndExitIfError(cmdStr)
}
