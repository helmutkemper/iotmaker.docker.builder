package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// ContainerInspect
//
// English: inspects the container
//
// Português: inspeciona o container
func (e *ContainerBuilder) ContainerInspect() (inspect iotmakerdocker.ContainerInspect, err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	inspect, err = e.dockerSys.ContainerInspectParsed(e.containerID)
	if err != nil {
		util.TraceToLog()
	}
	return
}
