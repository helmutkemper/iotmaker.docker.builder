package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"time"
)

// ContainerStartAfterBuild (english):
//
// ContainerStartAfterBuild (português): Inicia um container recem criado.
//
//   Saída:
//     err: Objeto de erro padrão
//
//   Nota: - Ha duas formas de criar um container:
//           ContainerBuildAndStartFromImage, inicializa o oncontainer e inicializa o
//           registro aa rede docker, para que o mesmo funcione de forma correta.
//           ContainerBuildWithoutStartingItFromImage apenas cria o container, por isto, a
//           primeira vez que o mesmo roda, ele deve ter o seu registro de rede
//           inicializado para que possa funcionar de forma correta.
//         - Apos inicializado a primeira vez, use as funções, ContainerStart,
//           ContainerPause e ContainerStop, caso necessite controlar o container.
func (e *ContainerBuilder) ContainerStartAfterBuild() (err error) {
	err = e.dockerSys.ContainerStart(e.containerID)
	if err != nil {
		util.TraceToLog()
		return
	}

	if e.waitString != "" && e.waitStringTimeout == 0 {
		_, err = e.dockerSys.ContainerLogsWaitText(e.containerID, e.waitString, log.Writer())
		if err != nil {
			util.TraceToLog()
			return
		}

	} else if e.waitString != "" {
		_, err = e.dockerSys.ContainerLogsWaitTextWithTimeout(e.containerID, e.waitString, e.waitStringTimeout, log.Writer())
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	if e.network == nil {
		e.IPV4Address, err = e.FindCurrentIPV4Address()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	e.chaos.serviceStartedAt = time.Now()
	e.startedAfterBuild = true
	if len(*e.onContainerReady) == 0 {
		*e.onContainerReady <- true
	}
	return
}
