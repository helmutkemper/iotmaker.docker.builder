package iotmaker_docker_builder

import (
	"errors"
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"time"
)

// ContainerBuilder (english):
//
// ContainerBuilder (português): Gerenciador de containers e imagens docker
type ContainerBuilder struct {
	network            isolatedNetwork.ContainerBuilderNetworkInterface
	dockerSys          iotmakerdocker.DockerSystem
	changePointer      *chan iotmakerdocker.ContainerPullStatusSendToChannel
	onContainerReady   *chan bool
	onContainerInspect *chan bool
	imageName          string
	imageID            string
	containerName      string
	buildPath          string
	environmentVar     []string
	changePorts        []changePort
	openPorts          []string
	doNotOpenPorts     bool
	waitString         string
	waitStringTimeout  time.Duration
	containerID        string
	ticker             *time.Ticker
	inspect            iotmakerdocker.ContainerInspect
	logs               string
	inspectInterval    time.Duration
	gitData            gitData
}

type Overloading struct {
	builder  *ContainerBuilder
	overload *ContainerBuilder
}

func (e Overloading) Init() (err error) {
	if e.builder == nil {
		err = errors.New("set container builder pointer first")
		return
	}

	e.overload = &ContainerBuilder{}
	e.overload.network = e.builder.network
	e.overload.SetImageName("overload:latest")
	e.overload.SetContainerName(e.builder.containerName + "_overload")
	e.overload.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.network.util.overload.image.git")
	e.overload.SetWaitString("overloading...")
	e.overload.SetEnvironmentVar(
		[]string{
			`IN_ADDRESS=10.0.0.6:8080`,
			`OUT_ADDRESS=10.0.0.3:8080`,
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
