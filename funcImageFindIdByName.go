package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) ImageFindIdByName(name string) (id string, err error) {
	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	id, err = e.dockerSys.ImageFindIdByName(name)
	if err != nil && err.Error() != "image name not found" {
		util.TraceToLog()
	}
	return
}
