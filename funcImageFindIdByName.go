package iotmakerdockerbuilder

import (
	dockerfileGolang "github.com/helmutkemper/iotmaker.docker.builder.golang.dockerfile"
)

func (e *ContainerBuilder) ImageFindIdByName(name string) (id string, err error) {
	if e.autoDockerfile == nil {
		e.autoDockerfile = &dockerfileGolang.DockerfileGolang{}
	}

	return e.dockerSys.ImageFindIdByName(name)
}
