package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// ContainerFindIdByName
//
// Similar:
//
//   ContainerFindIdByName(), ContainerFindIdByNameContains()
//
// English:
//
//  Searches and returns the ID of the container, if it exists
//
//   Input:
//     name: Full name of the container.
//
//   Output:
//     id: container ID
//     err: standard error object
//
// Português:
//
//  Procura e retorna o ID do container, caso o mesmo exista
//
//   Entrada:
//     name: Nome completo do container.
//
//   Saída:
//     id: ID do container
//     err: Objeto de erro padrão
func (e *ContainerBuilder) ContainerFindIdByName(name string) (id string, err error) {
	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	id, err = e.dockerSys.ContainerFindIdByName(name)
	if err != nil {
		util.TraceToLog()
	}

	return
}
