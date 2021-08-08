package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) ContainerFindIdByName(name string) (id string, err error) {
	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	id, err = e.dockerSys.ContainerFindIdByName(name)
	if err != nil {
		util.TraceToLog()
	}

	return
}
