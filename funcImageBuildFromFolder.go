package iotmakerdockerbuilder

import (
	"errors"
	"github.com/docker/docker/api/types"
	"path/filepath"
)

// ImageBuildFromFolder (english):
//
// ImageBuildFromFolder (português): transforma o conteúdo da pasta definida em SetBuildFolderPath() em uma imagem
//
//     Nota: A pasta deve conter um arquivo dockerfile, mas, como diferentes usos podem ter diferentes dockerfiles,
//     será dada a seguinte ordem na busca pelo arquivo: "Dockerfile-iotmaker", "Dockerfile", "dockerfile" na pasta raiz.
//     Se não houver encontrado, será feita uma busca recusiva por "Dockerfile" e "dockerfile"
//
func (e *ContainerBuilder) ImageBuildFromFolder() (err error) {
	err = e.verifyImageName()
	if err != nil {
		return
	}

	if e.buildPath == "" {
		err = errors.New("set build folder path first")
		return
	}

	e.buildPath, err = filepath.Abs(e.buildPath)
	if err != nil {
		return
	}

	var buildOptions types.ImageBuildOptions
	if buildOptions.BuildArgs == nil {
		buildOptions.BuildArgs = make(map[string]*string)
	}

	if e.contentGitConfigFile != "" {
		buildOptions.BuildArgs["GITCONFIG_FILE"] = &e.contentGitConfigFile
	}

	if e.contentIdRsaFile != "" {
		buildOptions.BuildArgs["SSH_ID_RSA_FILE"] = &e.contentIdRsaFile
	}

	e.imageID, err = e.dockerSys.ImageBuildFromFolder(
		e.buildPath,
		e.imageName,
		[]string{},
		buildOptions,
		e.changePointer,
	)
	if err != nil {
		err = errors.New(err.Error() + "\nfolder path: " + e.buildPath)
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
