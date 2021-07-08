package iotmakerdockerbuilder

import (
	dockerfileGolang "github.com/helmutkemper/iotmaker.docker.builder.golang.dockerfile"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

func (e *ContainerBuilder) ContainerFindIdByNameContains(containsName string) (list []NameAndId, err error) {
	list = make([]NameAndId, 0)

	if e.autoDockerfile == nil {
		e.autoDockerfile = &dockerfileGolang.DockerfileGolang{}
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
