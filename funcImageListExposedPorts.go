package iotmaker_docker_builder

import (
	"github.com/docker/go-connections/nat"
)

func (e *ContainerBuilder) ImageListExposedPorts() (list []nat.Port, err error) {

	list, err = e.dockerSys.ImageListExposedPortsByName(e.imageName)
	return
}
