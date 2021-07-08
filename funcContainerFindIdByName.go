package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

func (e *ContainerBuilder) ContainerFindIdByName(name string) (id string, err error) {
	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		return
	}

	return e.dockerSys.ContainerFindIdByName(name)
}
