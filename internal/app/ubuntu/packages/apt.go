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
	"xsel"}

func installAptPackages() {
	removeAptRepositories()
	addAptRepositories()

	packages := strings.Join(aptPackages[:], " ")

	exec.Command("sudo", "apt-get", "update")
	cmd := exec.Command("sudo", "apt-get", "-y", "install", packages)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%s", output)

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

	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()

	helpers.LogAndExitIfError(err)
}

func addYarnAptRepository() {
	const yarnPublicKeyURL = "https://dl.yarnpkg.com/debian/pubkey.gpg"
	const echoYarnList = `echo "deb https://dl.yarnpkg.com/debian/ stable main"`

	cmdStr := fmt.Sprintf(`
		curl -sS %s | sudo apt-key add - && \
		%s | sudo tee /etc/apt/sources.list.d/yarn.list`,
		yarnPublicKeyURL,
		echoYarnList)

	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()

	helpers.LogAndExitIfError(err)
}

func addDockerAptRepository() {
	const dockerPublicKeyURL = "https://download.docker.com/linux/ubuntu/gpg"
	const echoDockerList = `echo "deb http://apt.kubernetes.io/ kubernetes-xenial main"`

	cmdStr := fmt.Sprintf(`
		curl -fsSL %s | sudo apt-key add - && \
		 %s | sudo tee /etc/apt/sources.list.d/kubernetes.list`,
		dockerPublicKeyURL,
		echoDockerList)

	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()

	helpers.LogAndExitIfError(err)
}

func addKubectlAptRepository() {
	const kubectlPublicKeyURL = "https://packages.cloud.google.com/apt/doc/apt-key.gpg"
	const echoKubectlList = `echo "deb http://apt.kubernetes.io/ kubernetes-xenial main"`

	cmdStr := fmt.Sprintf(`
		curl -s %s | sudo apt-key add - && \
		 %s | sudo tee /etc/apt/sources.list.d/kubernetes.list`,
		kubectlPublicKeyURL,
		echoKubectlList)

	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()

	helpers.LogAndExitIfError(err)
}

func postInstallDocker() {
	const cmdStr = `
		sudo groupadd docker && \
		sudo usermod -aG docker "${USER}"`

	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()

	helpers.LogAndExitIfError(err)
}
