package iotmakerdockerbuilder

import (
	"os"
)

// SetDockerfilePath
//
// English:
//
// Defines a Dockerfile to build the image.
//
// Português:
//
// Define um arquivo Dockerfile para construir a imagem.
func (e *ContainerBuilder) SetDockerfilePath(path string) (err error) {
	var data []byte
	data, err = os.ReadFile(path)
	if err != nil {
		return err
	}

	e.buildOptions.Dockerfile = string(data)
	return
}
