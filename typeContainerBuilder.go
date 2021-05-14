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
	serverBuildPath    string
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
}

// ImageRemove (english):
//
// ImageRemove (português): remove a imagem se não houver containers usando a imagem (remova todos os containers antes
// do uso, mesmo os parados)
func (e *ContainerBuilder) ImageRemove() (err error) {
	err = e.dockerSys.ImageRemoveByName(e.imageName, false, false)
	return
}

// ContainerInspect (english):
//
// ContainerInspect (português): inspeciona o container
func (e *ContainerBuilder) ContainerInspect() (inspect iotmakerdocker.ContainerInspect, err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			return
		}
	}

	inspect, err = e.dockerSys.ContainerInspectParsed(e.containerID)
	return
}

// GetIdByContainerName (english):
//
// GetIdByContainerName (português): retorna o ID do container definido em SetContainerName()
func (e *ContainerBuilder) GetIdByContainerName() (err error) {
	e.containerID, err = e.dockerSys.ContainerFindIdByName(e.containerName)
	return
}

// RemoveAllByNameContains (english):
//
// RemoveAllByNameContains (português): procuta por redes, volumes, container e imagens que contenham o termo definido
// em "value" no nome e tenta remover os mesmos
func (e *ContainerBuilder) RemoveAllByNameContains(value string) (err error) {
	e.containerID = ""
	err = e.dockerSys.RemoveAllByNameContains(value)
	return
}
