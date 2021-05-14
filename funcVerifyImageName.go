package iotmaker_docker_builder

import (
	"errors"
	"strings"
)

// verifyImageName (english):
//
// verifyImageName (português): verifica se o nome da imagem tem a tag de versão
func (e *ContainerBuilder) verifyImageName() (err error) {
	if e.imageName == "" {
		err = errors.New("image name is't set")
		return
	}

	if strings.Contains(e.imageName, ":") == false {
		err = errors.New("image name must have a tag version. example: image_name:latest")
		return
	}

	return
}
