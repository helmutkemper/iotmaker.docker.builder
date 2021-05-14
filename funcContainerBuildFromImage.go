package iotmaker_docker_builder

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"log"
)

// ContainerBuildFromImage (english):
//
// ContainerBuildFromImage (português): transforma uma imagem baixada por ImagePull() ou criada por
// ImageBuildFromFolder() em container
func (e *ContainerBuilder) ContainerBuildFromImage() (err error) {
	err = e.verifyImageName()
	if err != nil {
		return
	}

	_, err = e.dockerSys.ImageFindIdByName(e.imageName)
	if err != nil {
		return
	}

	var netConfig *network.NetworkingConfig
	if e.network != nil {
		netConfig, err = e.network.GetConfiguration()
		if err != nil {
			return
		}
	}

	var portMap = nat.PortMap{}
	var list []nat.Port
	list, err = e.dockerSys.ImageListExposedPortsByName(e.imageName)
	if err != nil {
		return
	}

	if e.doNotOpenPorts == true {
		portMap = nil
	} else if e.openPorts != nil {
		var port nat.Port
		for _, portToOpen := range e.openPorts {
			port, err = nat.NewPort("tcp", portToOpen)
			if err != nil {
				return
			}

			portMap[port] = []nat.PortBinding{{HostPort: port.Port()}}
		}
	} else if e.changePorts != nil {
		var imagePort nat.Port
		var newPort nat.Port

		for _, newPortLinkMap := range e.changePorts {
			imagePort, err = nat.NewPort("tcp", newPortLinkMap.oldPort)
			if err != nil {
				return
			}

			newPort, err = nat.NewPort("tcp", newPortLinkMap.newPort)
			if err != nil {
				return
			}
			portMap[imagePort] = []nat.PortBinding{{HostPort: newPort.Port()}}
		}

	} else {
		for _, port := range list {
			portMap[port] = []nat.PortBinding{{HostPort: port.Port()}}
		}
	}

	var config = container.Config{
		OpenStdin:    true,
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Env:          []string{},
		Image:        e.imageName,
	}

	e.containerID, err = e.dockerSys.ContainerCreateWithConfig(
		&config,
		e.containerName,
		iotmakerdocker.KRestartPolicyNo,
		portMap,
		nil,
		netConfig,
	)
	if err != nil {
		return
	}

	err = e.dockerSys.ContainerStart(e.containerID)
	if err != nil {
		return
	}

	if e.waitString != "" {
		_, err = e.dockerSys.ContainerLogsWaitText(e.containerID, e.waitString, log.Writer())
		if err != nil {
			return
		}
	}

	*e.onContainerReady <- true
	return
}
