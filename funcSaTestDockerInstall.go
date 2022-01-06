package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// SaTestDockerInstall
//
// English:
//
//  Test if docker is responding correctly
//
//   Output:
//     err: Standard error object
//
// Português:
//
//  Testa se o docker está respondendo de forma correta
//
//   Saída:
//     err: Standard error object
func SaTestDockerInstall() (err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = dockerSys.ImageList()
	return
}
