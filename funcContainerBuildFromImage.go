package iotmakerdockerbuilder

import (
	"errors"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"log"
	"strings"
)

// ContainerBuildAndStartFromImage
//
// English: transforms an image downloaded by ImagePull() or created by ImageBuildFromFolder() into a container and start it
//
// Português: transforma uma imagem baixada por ImagePull() ou criada por ImageBuildFromFolder() em container e o inicializa
func (e *ContainerBuilder) ContainerBuildAndStartFromImage() (err error) {
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
		e.IPV4Address, err = e.FindCurrentIPV4Address()
	}

	*e.onContainerReady <- true
	return
}

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
		e.IPV4Address, netConfig, err = e.network.GetConfiguration()
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
			imagePort, err = nat.NewPort("tcp", newPortLinkMap.OldPort)
			if err != nil {
				return
			}

			var pass = false
			for _, portToVerify := range originalImagePortlist {
				if portToVerify.Port() == newPortLinkMap.OldPort {
					pass = true
					break
				}
			}

			if pass == false {
				err = errors.New("port " + newPortLinkMap.OldPort + " not found in image port list. port list: " + originalImagePortlistAsString)
				return
			}

			newPort, err = nat.NewPort("tcp", newPortLinkMap.NewPort)
			if err != nil {
				return
			}
			portMap[imagePort] = []nat.PortBinding{{HostPort: newPort.Port()}}
		}
	}

	if e.printBuildOutput == true {
		go func(ch *chan iotmakerdocker.ContainerPullStatusSendToChannel) {
			for {

				select {
				case event := <-*ch:
					var stream = event.Stream
					stream = strings.ReplaceAll(stream, "\n", "")
					stream = strings.ReplaceAll(stream, "\r", "")
					stream = strings.Trim(stream, " ")

					if stream == "" {
						continue
					}

					log.Printf("%v", stream)

					if event.Closed == true {
						return
					}
				}
			}
		}(&e.changePointer)
	}

	e.containerConfig.OpenStdin = true
	e.containerConfig.AttachStderr = true
	e.containerConfig.AttachStdin = true
	e.containerConfig.AttachStdout = true
	e.containerConfig.Env = e.environmentVar
	e.containerConfig.Image = e.imageName

	e.containerID, err = e.dockerSys.ContainerCreateWithConfig(
		&e.containerConfig,
		e.containerName,
		e.restartPolicy,
		portMap,
		e.volumes,
		netConfig,
	)
	if err != nil {
		return
	}

	*e.onContainerReady <- false
	return
}
