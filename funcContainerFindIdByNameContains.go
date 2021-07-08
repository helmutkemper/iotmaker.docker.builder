package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

func (e *ContainerBuilder) ContainerFindIdByNameContains(containsName string) (list []NameAndId, err error) {
	list = make([]NameAndId, 0)

	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		return
	}

	var recevedLis []iotmakerdocker.NameAndId
	recevedLis, err = e.dockerSys.ContainerFindIdByNameContains(containsName)
	if err != nil {
		return
	}

	for _, elementInList := range recevedLis {
		list = append(list, NameAndId(elementInList))
	}

	return
}
