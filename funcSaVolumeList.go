package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// SaVolumeList
//
// English:
//
//  List all docker volumes
//
//   Output:
//     list: docker volume list
//     err: Standard error object
//
// Português:
//
//  Lista todos os volumes docker
//
//   Saída:
//     list: lista de volumes docker
//     err: Objeto de erro padrão
func SaVolumeList() (list []types.Volume, err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	list, err = dockerSys.VolumeList()
	return
}
