package iotmakerdockerbuilder

import (
	"errors"
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	"github.com/helmutkemper/util"
	"log"
	"strconv"
	"time"
)

type NetworkChaos struct {
	imageName     string
	builder       *ContainerBuilder
	overload      *ContainerBuilder
	network       isolatedNetwork.ContainerBuilderNetworkInterface
	containerName string
	listenPort    int
	outputPort    int
	invert        bool
}

// SetNetworkDocker (english):
//
// SetNetworkDocker (português): Define o ponteiro do gerenciador de rede docker
//   Entrada:
//     network: ponteiro para o objeto gerenciador de rede.
//
//   Nota: - A entrada network deve ser compatível com a interface
//           dockerBuilderNetwork.ContainerBuilderNetwork{}
func (e *NetworkChaos) SetNetworkDocker(network isolatedNetwork.ContainerBuilderNetworkInterface) {
	e.network = network
}

func (e *NetworkChaos) SetFatherContainer(fatherContainer *ContainerBuilder) {
	e.builder = fatherContainer
}

func (e *NetworkChaos) SetContainerName(value string) {
	e.containerName = value
}

func (e *NetworkChaos) SetPorts(listenPort, outputPort int, invert bool) {
	e.listenPort = listenPort
	e.outputPort = outputPort
	e.invert = invert
}

func (e *NetworkChaos) Init() (err error) {
	if e.builder == nil {
		err = errors.New("father container must be set")
		return
	}

	if e.containerName == "" {
		//err = errors.New("containerName must be set")
		//return
	}

	if e.listenPort == 0 {
		err = errors.New("listen port must be set")
		return
	}

	if e.outputPort == 0 {
		err = errors.New("output port must be set")
		return
	}

	var builderIP = e.builder.GetIPV4Address()
	var nextIP, _ = e.builder.incIpV4Address(builderIP, 1)

	e.imageName = "overload."

	if e.invert == false {

		e.imageName += "in.port." +
			strconv.FormatInt(int64(e.outputPort), 10) +
			".out.port." +
			strconv.FormatInt(int64(e.listenPort), 10)

	} else {

		e.imageName += "in.port." +
			strconv.FormatInt(int64(e.listenPort), 10) +
			".out.port." +
			strconv.FormatInt(int64(e.outputPort), 10)

	}

	e.imageName += ":latest"

	e.overload = &ContainerBuilder{}
	e.overload.network = e.builder.network
	e.overload.SetImageName(e.imageName)
	e.overload.MakeDefaultDockerfileForMe()
	log.Printf("container name: %v", e.builder.containerName+"_overload")
	e.overload.SetContainerName(e.builder.containerName + "_overload")
	e.overload.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.network.util.overload.image.git")
	e.overload.SetWaitStringWithTimeout("overloading...", 10*time.Second)
	e.overload.AddPortToExpose("27016")
	e.overload.AddPortToExpose("8080")
	//e.overload.AddPortToExpose("8080")
	if e.invert == false {
		log.Printf(`1.IN_ADDRESS=` + nextIP + `:` + strconv.FormatInt(int64(e.outputPort), 10))
		log.Printf(`1.OUT_ADDRESS=` + builderIP + `:` + strconv.FormatInt(int64(e.listenPort), 10))
		e.overload.SetEnvironmentVar(
			[]string{
				`IN_ADDRESS=` + nextIP + `:` + strconv.FormatInt(int64(e.outputPort), 10),
				`OUT_ADDRESS=` + builderIP + `:` + strconv.FormatInt(int64(e.listenPort), 10),
				`MIN_DELAY=10`,
				`MAX_DELAY=11`,
			},
		)

	} else {
		log.Printf(`2.OUT_ADDRESS=` + nextIP + `:` + strconv.FormatInt(int64(e.outputPort), 10))
		log.Printf(`2.IN_ADDRESS=` + builderIP + `:` + strconv.FormatInt(int64(e.listenPort), 10))
		e.overload.SetEnvironmentVar(
			[]string{
				`OUT_ADDRESS=` + nextIP + `:` + strconv.FormatInt(int64(e.outputPort), 10),
				`IN_ADDRESS=` + builderIP + `:` + strconv.FormatInt(int64(e.listenPort), 10),
				`MIN_DELAY=10`,
				`MAX_DELAY=11`,
			},
		)
	}

	err = e.overload.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	if e.imageExists() == false {
		err = e.overload.ImageBuildFromServer()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.overload.ContainerBuildFromImage()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

func (e *NetworkChaos) imageExists() (exists bool) {
	var ID string
	ID, _ = e.overload.dockerSys.ImageFindIdByName(e.imageName)
	exists = ID != ""
	return
}
