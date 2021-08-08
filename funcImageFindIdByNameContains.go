package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) ImageFindIdByNameContains(containsName string) (list []NameAndId, err error) {
	list = make([]NameAndId, 0)

	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	var recevedLis []iotmakerdocker.NameAndId
	recevedLis, err = e.dockerSys.ImageFindIdByNameContains(containsName)
	if err != nil {
		util.TraceToLog()
		return
	}

	for _, elementInList := range recevedLis {
		list = append(list, NameAndId(elementInList))
	}

	return
}
