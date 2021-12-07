package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"io"
)

func (e *ContainerBuilder) ContainerCopyTo(
	destinationPath string,
	content io.Reader,
) (
	err error,
) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerCopyTo(e.containerID, destinationPath, content)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
