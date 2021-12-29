package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// ContainerStart
//
// English:
//
//  Initialize a paused or stoped container
//
//   Output:
//     err: Default error object.
//
// Note:
//
//   * There are two ways to create a container:
//     ContainerBuildAndStartFromImage, initializes the oncontainer and initializes the registry to
//     the docker network, so that it works correctly.
//     ContainerBuildWithoutStartingItFromImage just creates the container, so the first time it runs,
//     it must have its network registry initialized so it can work properly.
//   * After initializing the first time, use the functions, ContainerStart, ContainerPause and
//     ContainerStop, if you need to control the container.
//
// Português:
//
//  Inicializar um container pausado ou parado.
//
//   Saída:
//     err: Objeto de erro padrão.
//
// Nota:
//
//   * Ha duas formas de criar um container:
//     ContainerBuildAndStartFromImage, inicializa o oncontainer e inicializa o registro aa rede
//     docker, para que o mesmo funcione de forma correta.
//     ContainerBuildWithoutStartingItFromImage apenas cria o container, por isto, a primeira vez que
//     o mesmo roda, ele deve ter o seu registro de rede inicializado para que possa funcionar de
//     forma correta.
//   * Apos inicializado a primeira vez, use as funções, ContainerStart, ContainerPause e
//     ContainerStop, caso necessite controlar o container.
func (e *ContainerBuilder) ContainerStart() (err error) {

	if e.startedAfterBuild == false {
		err = e.ContainerStartAfterBuild()
		if err != nil {
			util.TraceToLog()
			return
		}
		return
	}

	if e.containerID == "" {
		err = e.getIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerStart(e.containerID)
	if err != nil {
		util.TraceToLog()
	}
	return
}
