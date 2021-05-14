package iotmaker_docker_builder

import (
	"errors"
	"path/filepath"
)

func (e *ContainerBuilder) ImageBuildFromServer() (err error) {
	err = e.verifyImageName()
	if err != nil {
		return
	}

	if e.serverBuildPath == "" {
		err = errors.New("set server url to build first")
		return
	}

	e.buildPath, err = filepath.Abs(e.buildPath)
	if err != nil {
		return
	}

	e.imageID, err = e.dockerSys.ImageBuildFromRemoteServer(
		e.serverBuildPath,
		e.imageName,
		[]string{},
		e.changePointer,
	)
	if err != nil {
		return
	}

	if e.imageID == "" {
		err = errors.New("image ID was not generated")
		return
	}

	// Construir uma imagem de múltiplas etapas deixa imagens grandes e sem serventia, ocupando espaço no HD.
	err = e.dockerSys.ImageGarbageCollector()
	if err != nil {
		return
	}

	return
}
