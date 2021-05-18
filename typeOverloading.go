package iotmaker_docker_builder

import (
	"errors"
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	"log"
	"time"
)

type Overloading struct {
	builder  *ContainerBuilder
	overload *ContainerBuilder
	network  isolatedNetwork.ContainerBuilderNetworkInterface
}

//fixme: interface
func (e *Overloading) SetBuilderToOverload(value *ContainerBuilder) {
	e.builder = value
}

// SetNetworkDocker (english):
//
// SetNetworkDocker (português): Define o ponteiro do gerenciador de rede docker
//   network: ponteiro para o objeto gerenciador de rede.
//
//     Nota: compatível com o objeto dockerBuilderNetwork.ContainerBuilderNetwork{}
func (e *Overloading) SetNetworkDocker(network isolatedNetwork.ContainerBuilderNetworkInterface) {
	e.network = network
}

func (e *Overloading) Init(listenPort string, invert bool) (err error) {
	if e.builder == nil {
		err = errors.New("set container builder pointer first")
		return
	}

	var builderIP = e.builder.GetIPV4Address()
	var nextIP, _ = e.builder.incIpV4Address(builderIP, 1)

	e.overload = &ContainerBuilder{}
	e.overload.network = e.builder.network
	e.overload.SetImageName("overload:latest")
	e.overload.SetContainerName(e.builder.containerName + "_overload")
	e.overload.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.network.util.overload.image.git")
	e.overload.SetWaitStringWithTimeout("overloading...", 10*time.Second)
	e.overload.AddPortToOpen("8000")
	e.overload.AddPortToOpen("8080")
	if invert == false {
		log.Printf(`1.IN_ADDRESS=` + nextIP + `:8000`)
		log.Printf(`1.OUT_ADDRESS=` + builderIP + `:` + listenPort)
		e.overload.SetEnvironmentVar(
			[]string{
				`IN_ADDRESS=` + nextIP + `:8000`,
				`OUT_ADDRESS=` + builderIP + `:` + listenPort,
				`MIN_DELAY=10`,
				`MAX_DELAY=11`,
			},
		)
	} else {
		log.Printf(`2.OUT_ADDRESS=` + nextIP + `:8000`)
		log.Printf(`2.IN_ADDRESS=` + builderIP + `:` + listenPort)
		e.overload.SetEnvironmentVar(
			[]string{
				`OUT_ADDRESS=` + nextIP + `:8000`,
				`IN_ADDRESS=` + builderIP + `:` + listenPort,
				`MIN_DELAY=10`,
				`MAX_DELAY=11`,
			},
		)
	}

	err = e.overload.Init()
	if err != nil {
		return
	}

	err = e.overload.ImageBuildFromServer()
	if err != nil {
		return
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
