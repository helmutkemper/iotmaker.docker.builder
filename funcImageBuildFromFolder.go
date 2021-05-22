package iotmakerdockerbuilder

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
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

	if e.makeDefaultDockerfile == true {
		var dockerfile string
		var fileList []fs.FileInfo

		fileList, err = ioutil.ReadDir(e.buildPath)
		if err != nil {
			return
		}

		var pass = false
		for _, file := range fileList {
			if file.Name() == "go.mod" {
				pass = true
				break
			}
		}
		if pass == false {
			err = errors.New("go.mod file not found")
			return
		}

		dockerfile, err = e.mountDefaultGolangDockerfile()
		if err != nil {
			return
		}

		var dockerfilePath = filepath.Join(e.buildPath, "Dockerfile-iotmaker")
		err = ioutil.WriteFile(dockerfilePath, []byte(dockerfile), os.ModePerm)
		if err != nil {
			return
		}
	}

	e.imageID, err = e.dockerSys.ImageBuildFromFolder(
		e.buildPath,
		e.imageName,
		[]string{},
		e.buildOptions,
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
