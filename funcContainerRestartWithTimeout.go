package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"time"
)

func (e *ContainerBuilder) ContainerRestartWithTimeout(timeout time.Duration) (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerRestartWithTimeout(e.containerID, timeout)
	if err != nil {
		util.TraceToLog()
	}
	return
}
