package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// SaVolumeCreate
//
// English:
//
//  Create a docker volume
//
//   Input:
//     labels: labels list
//     name: volume name
//
//   Output:
//     list: docker volume list
//     err: Standard error object
//
// Português:
//
//  Cria um volume docker
//
//   Input:
//     labels: lista de rótulos
//     name: nome do volume
//
//   Saída:
//     list: lista de volumes docker
//     err: Objeto de erro padrão
func SaVolumeCreate(
	labels map[string]string,
	name string,
) (
	volume types.Volume,
	err error,
) {

	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	volume, err = dockerSys.VolumeCreate(labels, name)
	return
}
