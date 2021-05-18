package iotmaker_docker_builder

import (
	"errors"
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
		e.ipAddress, netConfig, err = e.network.GetConfiguration()
		if err != nil {
			return
		}
	}

	var portMap = nat.PortMap{}
	var originalImagePortlist []nat.Port
	var originalImagePortlistAsString string
	originalImagePortlist, err = e.dockerSys.ImageListExposedPortsByName(e.imageName)

	if err != nil {
		return
	}

	for k, v := range originalImagePortlist {
		if k != 0 {
			originalImagePortlistAsString += ", "
		}
		originalImagePortlistAsString += v.Port()
	}

	if e.openAllPorts == true {
		for _, port := range originalImagePortlist {
			portMap[port] = []nat.PortBinding{{HostPort: port.Port()}}
		}
	} else if e.openPorts != nil {
		var port nat.Port
		for _, portToOpen := range e.openPorts {
			var pass = false
			for _, portToVerify := range originalImagePortlist {
				if portToVerify.Port() == portToOpen {
					pass = true
					break
				}
			}

			if pass == false {
				err = errors.New("port " + portToOpen + " not found in image port list. port list: " + originalImagePortlistAsString)
				return
			}

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

			var pass = false
			for _, portToVerify := range originalImagePortlist {
				if portToVerify.Port() == newPortLinkMap.oldPort {
					pass = true
					break
				}
			}

			if pass == false {
				err = errors.New("port " + newPortLinkMap.oldPort + " not found in image port list. port list: " + originalImagePortlistAsString)
				return
			}

			newPort, err = nat.NewPort("tcp", newPortLinkMap.newPort)
			if err != nil {
				return
			}
			portMap[imagePort] = []nat.PortBinding{{HostPort: newPort.Port()}}
		}
	}

	var config = container.Config{
		OpenStdin:    true,
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Env:          e.environmentVar,
		Image:        e.imageName,
	}

	e.containerID, err = e.dockerSys.ContainerCreateWithConfig(
		&config,
		e.containerName,
		iotmakerdocker.KRestartPolicyNo,
		portMap,
		e.volumes,
		netConfig,
	)
	if err != nil {
		return
	}

	err = e.dockerSys.ContainerStart(e.containerID)
	if err != nil {
		return
	}

	if e.waitString != "" && e.waitStringTimeout == 0 {
		_, err = e.dockerSys.ContainerLogsWaitText(e.containerID, e.waitString, log.Writer())
		if err != nil {
			return
		}
	} else if e.waitString != "" {
		_, err = e.dockerSys.ContainerLogsWaitTextWithTimeout(e.containerID, e.waitString, e.waitStringTimeout, log.Writer())
		if err != nil {
			return
		}
	}

	if e.network == nil {
		e.ipAddress, err = e.FindCurrentIpAddress()
	}

	*e.onContainerReady <- true
	return
}
