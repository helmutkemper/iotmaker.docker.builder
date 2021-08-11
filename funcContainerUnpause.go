package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) ContainerUnpause() (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerUnpause(e.containerID)
	if err != nil {
		util.TraceToLog()
	}
	return
}
