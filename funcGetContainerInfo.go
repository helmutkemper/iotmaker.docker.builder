package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
)

func (e *ContainerBuilder) GetContainerInfo() (info types.Info, err error) {
	info, err = e.dockerSys.DockerInfo()
	return
}
