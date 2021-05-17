package iotmaker_docker_builder

import (
	"errors"
	"time"
)

type Overloading struct {
	builder  *ContainerBuilder
	overload *ContainerBuilder
}

//fixme: interface
func (e *Overloading) SetBuilderToOverload(value *ContainerBuilder) {
	e.builder = value
}

func (e *Overloading) Init() (err error) {
	if e.builder == nil {
		err = errors.New("set container builder pointer first")
		return
	}

	e.overload = &ContainerBuilder{}
	e.overload.network = e.builder.network
	e.overload.SetImageName("overload:latest")
	e.overload.SetContainerName(e.builder.containerName + "_overload")
	e.overload.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.network.util.overload.image.git")
	e.overload.SetWaitStringWithTimeout("overloading...", 10*time.Second)
	e.overload.SetEnvironmentVar(
		[]string{
			`IN_ADDRESS=10.0.0.3:3030`,
			`OUT_ADDRESS=10.0.0.2:3030`,
			`MIN_DELAY=100`,
			`MAX_DELAY=1500`,
		},
	)

	err = e.overload.Init()
	if err != nil {
		return
	}

	if e.imageExists() == false {
		err = e.overload.ImageBuildFromServer()
		if err != nil {
			return
		}
	}

	err = e.overload.ContainerBuildFromImage()
	if err != nil {
		return
	}

	return
}

func (e *Overloading) imageExists() (exists bool) {
	var ID string
	ID, _ = e.overload.dockerSys.ImageFindIdByName("overload:latest")
	exists = ID != ""
	return
}
