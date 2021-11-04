package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// ContainerRemove
//
// English: stop and remove the container
//   Input:
//     removeVolumes: removes docker volumes linked to the container
//   Output:
//     err: standard error object
//
// Português: parar e remover o container
//   Entrada:
//     removeVolumes: remove os volumes docker vinculados ao container
//   Saída:
//     err: Objeto de erro padrão
func (e *ContainerBuilder) ContainerRemove(removeVolumes bool) (err error) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerStopAndRemove(e.containerID, removeVolumes, false, false)
	if err != nil {
		util.TraceToLog()
	}
	return
}
