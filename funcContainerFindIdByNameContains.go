package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// ContainerFindIdByNameContains
//
// Similar:
//
//   ContainerFindIdByName(), ContainerFindIdByNameContains()
//
// English:
//
//  Searches and returns the ID list of the container name
//
//   Input:
//     name: name of the container.
//
//   Output:
//     id: list of containers ID
//     err: standard error object
//
// Português:
//
//  Procura e retorna uma lista de IDs de containers
//
//   Entrada:
//     name: Nome do container.
//
//   Saída:
//     id: lista de IDs dos containers
//     err: Objeto de erro padrão
func (e *ContainerBuilder) ContainerFindIdByNameContains(containsName string) (list []NameAndId, err error) {
	list = make([]NameAndId, 0)

	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	var recevedLis []iotmakerdocker.NameAndId
	recevedLis, err = e.dockerSys.ContainerFindIdByNameContains(containsName)
	if err != nil {
		util.TraceToLog()
		return
	}

	for _, elementInList := range recevedLis {
		list = append(list, NameAndId(elementInList))
	}

	return
}
