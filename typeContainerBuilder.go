package iotmaker_docker_builder

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"log"
	"path/filepath"
	"strings"
	"time"
)

// changePort (english):
//
// changePort (português):
type changePort struct {
	oldPort string
	newPort string
}

// ContainerBuilder (english):
//
// ContainerBuilder (português):
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
}

// GetLastInspect (english):
//
// GetLastInspect (português):
func (e *ContainerBuilder) GetLastInspect() (inspect iotmakerdocker.ContainerInspect) {
	return e.inspect
}

// GetLastLogs (english):
//
// GetLastLogs (português):
func (e *ContainerBuilder) GetLastLogs() (logs string) {
	return e.logs
}

// SetBuildFolderPath (english):
//
// SetBuildFolderPath (português):
func (e *ContainerBuilder) SetBuildFolderPath(value string) {
	e.buildPath = value
}

// SetImageName (english):
//
// SetImageName (português):
func (e *ContainerBuilder) SetImageName(value string) {
	e.imageName = value
}

// SetContainerName (english):
//
// SetContainerName (português):
func (e *ContainerBuilder) SetContainerName(value string) {
	e.containerName = value
}

// SetWaitString (english):
//
// SetWaitString (português):
func (e *ContainerBuilder) SetWaitString(value string) {
	e.waitString = value
}

// SetNetworkDocker (english):
//
// SetNetworkDocker (português):
func (e *ContainerBuilder) SetNetworkDocker(network isolatedNetwork.ContainerBuilderNetworkInterface) {
	e.network = network
}

// SetEnvironmentVar (english):
//
// SetEnvironmentVar (português):
func (e *ContainerBuilder) SetEnvironmentVar(value []string) {
	e.environmentVar = value
}

// AddPortToOpen (english):
//
// AddPortToOpen (português):
func (e *ContainerBuilder) AddPortToOpen(value string) {
	if e.openPorts == nil {
		e.openPorts = make([]string, 0)
	}

	e.openPorts = append(e.openPorts, value)
}

// AddPortToChange (english):
//
// AddPortToChange (português):
func (e *ContainerBuilder) AddPortToChange(imagePort string, newPort string) {
	if e.changePorts == nil {
		e.changePorts = make([]changePort, 0)
	}

	e.changePorts = append(
		e.changePorts,
		changePort{
			oldPort: imagePort,
			newPort: newPort,
		},
	)
}

// SetDoNotOpenContainersPorts (english):
//
// SetDoNotOpenContainersPorts (português):
func (e *ContainerBuilder) SetDoNotOpenContainersPorts() {
	e.doNotOpenPorts = true
}

// SetInspectInterval (english):
//
// SetInspectInterval (português):
func (e *ContainerBuilder) SetInspectInterval(value time.Duration) {
	e.inspectInterval = value
}

// Init (english):
//
// Init (português):
func (e *ContainerBuilder) Init() (err error) {
	if e.environmentVar == nil {
		e.environmentVar = make([]string, 0)
	}

	onStart := make(chan bool, 1)
	e.onContainerReady = &onStart

	onInspect := make(chan bool, 1)
	e.onContainerInspect = &onInspect

	e.changePointer = iotmakerdocker.NewImagePullStatusChannel()

	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		return
	}

	if e.inspectInterval != 0 {
		e.ticker = time.NewTicker(e.inspectInterval)

		go func(e *ContainerBuilder) {
			var err error
			var logs []byte

			for {
				select {
				case <-e.ticker.C:

					if e.containerID == "" {
						e.containerID, err = e.dockerSys.ContainerFindIdByName(e.containerName)
						if err != nil {
							continue
						}
					}

					e.inspect, _ = e.dockerSys.ContainerInspectParsed(e.containerID)
					logs, _ = e.dockerSys.ContainerLogs(e.containerID)
					e.logs = string(logs)
					*e.onContainerInspect <- true
				}
			}

		}(e)
	}

	return
}

// GetChannelOnContainerReady (english):
//
// GetChannelOnContainerReady (português):
func (e *ContainerBuilder) GetChannelOnContainerReady() (channel *chan bool) {
	return e.onContainerReady
}

// GetChannelOnContainerInspect (english):
//
// GetChannelOnContainerInspect (português):
func (e *ContainerBuilder) GetChannelOnContainerInspect() (channel *chan bool) {
	return e.onContainerInspect
}

// GetChannelEvent (english):
//
// GetChannelEvent (português):
func (e *ContainerBuilder) GetChannelEvent() (channel *chan iotmakerdocker.ContainerPullStatusSendToChannel) {
	return e.changePointer
}

// ImagePull (english):
//
// ImagePull (português):
func (e *ContainerBuilder) ImagePull() (err error) {
	e.imageID, e.imageName, err = e.dockerSys.ImagePull(e.imageName, e.changePointer)
	if err != nil {
		return
	}

	return
}

// verifyImageName (english):
//
// verifyImageName (português):
func (e *ContainerBuilder) verifyImageName() (err error) {
	if e.imageName == "" {
		err = errors.New("image name is't set")
		return
	}

	if strings.Contains(e.imageName, ":") == false {
		err = errors.New("image name must have a tag version. example: image_name:latest")
		return
	}

	return
}

// WaitFortextInContainerLog (english):
//
// WaitFortextInContainerLog (português):
func (e *ContainerBuilder) WaitFortextInContainerLog(value string) (dockerLogs string, err error) {
	var logs []byte
	logs, err = e.dockerSys.ContainerLogsWaitText(e.containerID, value, log.Writer())
	return string(logs), err
}

// ImageBuildFromFolder (english):
//
// ImageBuildFromFolder (português):
func (e *ContainerBuilder) ImageBuildFromFolder() (err error) {
	err = e.verifyImageName()
	if err != nil {
		return
	}

	e.buildPath, err = filepath.Abs(e.buildPath)
	if err != nil {
		return
	}

	e.imageID, err = e.dockerSys.ImageBuildFromFolder(
		e.buildPath,
		[]string{
			e.imageName,
		},
		e.changePointer,
	)
	if err != nil {
		return
	}

	if e.imageID == "" {
		err = errors.New("image ID was not generated")
		return
	}

	// Construir uma imagem de múltiplas etapas deixa imagens grandes e sem serventia, ocupando espaço no HD.
	err = e.dockerSys.ImageGarbageCollector()
	if err != nil {
		return
	}

	return
}

// ContainerBuildFromImage (english):
//
// ContainerBuildFromImage (português):
func (e *ContainerBuilder) ContainerBuildFromImage() (err error) {
	err = e.verifyImageName()
	if err != nil {
		return
	}

	//if e.network == nil {
	//  err = errors.New("network interface is't set")
	//  return
	//}

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

// GetContainerLog (english):
//
// GetContainerLog (português):
func (e *ContainerBuilder) GetContainerLog() (log []byte, err error) {
	if e.containerID == "" {
		err = e.GetFindIdByContainerName()
		if err != nil {
			return
		}
	}

	log, err = e.dockerSys.ContainerLogs(e.containerID)
	return
}

// FindTextInsideContainerLog (english):
//
// FindTextInsideContainerLog (português):
func (e *ContainerBuilder) FindTextInsideContainerLog(value string) (contains bool, err error) {
	var logs []byte
	logs, err = e.GetContainerLog()
	if err != nil {
		return
	}

	contains = bytes.Contains(logs, []byte(value))
	return
}

// ContainerStart (english):
//
// ContainerStart (português):
func (e *ContainerBuilder) ContainerStart() (err error) {
	if e.containerID == "" {
		err = e.GetFindIdByContainerName()
		if err != nil {
			return
		}
	}

	err = e.dockerSys.ContainerStart(e.containerID)
	return
}

// ContainerStop (english):
//
// ContainerStop (português):
func (e *ContainerBuilder) ContainerStop() (err error) {
	if e.containerID == "" {
		err = e.GetFindIdByContainerName()
		if err != nil {
			return
		}
	}

	err = e.dockerSys.ContainerStop(e.containerID)
	return
}

// ContainerRemove (english):
//
// ContainerRemove (português):
func (e *ContainerBuilder) ContainerRemove() (err error) {
	if e.containerID == "" {
		err = e.GetFindIdByContainerName()
		if err != nil {
			return
		}
	}

	err = e.dockerSys.ContainerStopAndRemove(e.containerID, true, false, false)
	return
}

// ImageRemove (english):
//
// ImageRemove (português):
func (e *ContainerBuilder) ImageRemove() (err error) {
	err = e.ContainerRemove()
	if err != nil {
		return
	}

	err = e.dockerSys.ImageRemoveByName(e.imageName, false, false)
	return
}

// ContainerInspect (english):
//
// ContainerInspect (português):
func (e *ContainerBuilder) ContainerInspect() (inspect iotmakerdocker.ContainerInspect, err error) {
	if e.containerID == "" {
		err = e.GetFindIdByContainerName()
		if err != nil {
			return
		}
	}

	inspect, err = e.dockerSys.ContainerInspectParsed(e.containerID)
	return
}

// GetFindIdByContainerName (english):
//
// GetFindIdByContainerName (português):
func (e *ContainerBuilder) GetFindIdByContainerName() (err error) {
	e.containerID, err = e.dockerSys.ContainerFindIdByName(e.containerName)
	return
}

// RemoveAllByNameContains (english):
//
// RemoveAllByNameContains (português):
func (e *ContainerBuilder) RemoveAllByNameContains(name string) (err error) {
	e.containerID = ""
	err = e.dockerSys.RemoveAllByNameContains(name)
	return
}
