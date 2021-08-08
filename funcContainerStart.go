package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// ContainerStart
//
// English: initialize a newly created or paused container
//
// Português: inicializar um container recém criado ou pausado
func (e *ContainerBuilder) ContainerStart() (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerStart(e.containerID)
	if err != nil {
		util.TraceToLog()
	}
	return
}
