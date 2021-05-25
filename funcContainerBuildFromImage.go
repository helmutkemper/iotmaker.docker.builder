package iotmakerdockerbuilder

import (
	"errors"
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

// SetContainerRestartPolicy (english):
//
// SetContainerRestartPolicy (português): Define a política de reinício do container.
//   value: KRestartPolicyNo            - não reiniciar o container (padrão).
//          KRestartPolicyOnFailure     - reinicia o container se houver um erro (com o manifesto informando um código de erro diferente de zero).
//          KRestartPolicyAlways        - sempre reinicia o container quando ele para, mesmo quando ele é parado manualmente.
//          KRestartPolicyUnlessStopped - sempre reinicia o container, caso ele não tenha sido parado manualmente.
func (e *ContainerBuilder) SetContainerRestartPolicy(value RestartPolicy) {
	e.restartPolicy = iotmakerdocker.RestartPolicy(value)
}

type RestartPolicy int

const (
	//KRestartPolicyNo (english): Do not automatically restart the container. (the default)
	KRestartPolicyNo RestartPolicy = iota

	//KRestartPolicyOnFailure (english): Restart the container if it exits due to an error, which manifests as a non-zero exit
	//code.
	KRestartPolicyOnFailure

	//KRestartPolicyAlways (english): Always restart the container if it stops. If it is manually stopped, it is restarted
	//only when Docker daemon restarts or the container itself is manually restarted. (See the second bullet listed in restart policy details)
	KRestartPolicyAlways

	//KRestartPolicyUnlessStopped (english): Similar to always, except that when the container is stopped (manually or otherwise),
	//it is not restarted even after Docker daemon restarts.
	KRestartPolicyUnlessStopped
)
