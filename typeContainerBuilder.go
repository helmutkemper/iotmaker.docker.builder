package iotmaker_docker_builder

import (
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
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
	doNotOpenPorts     bool
	waitString         string
	containerID        string
	ticker             *time.Ticker
	inspect            iotmakerdocker.ContainerInspect
	logs               string
	inspectInterval    time.Duration
	gitData            gitData
}

// gitData (Português): Estrutura de dados baseada no framework go-git
type gitData struct {
	url               string
	sshPrivateKeyPath string
	privateToke       string
	user              string
	password          string
}
