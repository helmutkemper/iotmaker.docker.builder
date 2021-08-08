package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// ContainerPause
//
// English: pause the container
//
// Português: pausa o container
func (e *ContainerBuilder) ContainerPause() (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerPause(e.containerID)
	if err != nil {
		util.TraceToLog()
	}
	return
}
