package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// ContainerInspect
//
// English:
//
//  Inspects the container
//
//   Output:
//     inspect: Contains information about the container, such as ID, name, status, volumes, etc.
//     err: Standard error object.
//
// Português:
//
//  Inspeciona o container
//
//   Saída:
//     inspect: Contém informações sobre o container, como ID, nome, status, volumes, etc.
//     err: Objeto de erro padrão.
func (e *ContainerBuilder) ContainerInspect() (inspect iotmakerdocker.ContainerInspect, err error) {
	if e.containerID == "" {
		err = e.getIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	inspect, err = e.dockerSys.ContainerInspectParsed(e.containerID)
	if err != nil {
		util.TraceToLog()
	}
	return
}
