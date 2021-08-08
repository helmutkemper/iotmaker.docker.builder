package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
)

// ImageRemoveByName
//
// English: remove the image if there are no containers using the image (remove all containers before use, including
// stopped containers)
//
//   name: full image name
//
// Português: remove a imagem se não houver containers usando a imagem (remova todos os containers antes
// do uso, inclusive os containers parados)
//
//   name: nome completo da imagem
func (e *ContainerBuilder) ImageRemoveByName(name string) (err error) {
	err = e.dockerSys.ImageRemoveByName(name, false, false)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
