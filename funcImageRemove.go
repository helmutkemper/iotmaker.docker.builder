package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// ImageRemove
//
// English: remove the image if there are no containers using the image (remove all containers before use, including
// stopped containers)
//
// Português: remove a imagem se não houver containers usando a imagem (remova todos os containers antes
// do uso, inclusive os containers parados)
func (e *ContainerBuilder) ImageRemove() (err error) {
	err = e.dockerSys.ImageRemoveByName(e.imageName, false, false)
	if err != nil {
		util.TraceToLog()
	}
	return
}
