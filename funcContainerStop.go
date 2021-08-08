package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// ContainerStop
//
// English: stop the container
//
// Português: parar o container
func (e *ContainerBuilder) ContainerStop() (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerStop(e.containerID)
	if err != nil {
		util.TraceToLog()
	}
	return
}
