package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
)

// TestDockerInstall
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
func (e ContainerBuilder) TestDockerInstall() (err error) {
	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = e.dockerSys.ImageList()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
