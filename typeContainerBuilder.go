package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types/mount"
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

// ContainerBuilder (english):
//
// ContainerBuilder (português): Gerenciador de containers e imagens docker
type ContainerBuilder struct {
	network            isolatedNetwork.ContainerBuilderNetworkInterface
	dockerSys          iotmakerdocker.DockerSystem
	changePointer      *chan iotmakerdocker.ContainerPullStatusSendToChannel
	onContainerReady   *chan bool
	onContainerInspect *chan bool
	imageName          string
	imageID            string
	containerName      string
	buildPath          string
	environmentVar     []string
	changePorts        []changePort
	openPorts          []string
	openAllPorts       bool
	waitString         string
	waitStringTimeout  time.Duration
	containerID        string
	ticker             *time.Ticker
	inspect            iotmakerdocker.ContainerInspect
	logs               string
	inspectInterval    time.Duration
	gitData            gitData
	volumes            []mount.Mount
	IPV4Address        string

	contentIdRsaFile      string
	contentKnownHostsFile string
	contentGitConfigFile  string
}

// SetPrivateRepositoryAutoConfig (english):
//
// SetPrivateRepositoryAutoConfig (português): Copia a chave ssh  ~/.ssh/id_rsa e o arquivo ~/.gitconfig para as
// variáveis SSH_ID_RSA_FILE e GITCONFIG_FILE.
//
// Dentro do Dockerfile
func (e *ContainerBuilder) SetPrivateRepositoryAutoConfig() (err error) {
	var userData *user.User
	var fileData []byte

	userData, err = user.Current()
	if err != nil {
		return
	}
	var filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "id_rsa")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}
	e.contentIdRsaFile = string(fileData)
	e.contentIdRsaFile = strings.ReplaceAll(e.contentIdRsaFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "known_hosts")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}
	e.contentKnownHostsFile = string(fileData)
	e.contentKnownHostsFile = strings.ReplaceAll(e.contentKnownHostsFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".gitconfig")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}
	e.contentGitConfigFile = string(fileData)
	e.contentGitConfigFile = strings.ReplaceAll(e.contentGitConfigFile, `"`, `\"`)
	return
}
