package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"io"
)

// ContainerCopyTo
//
// Português: Copia um arquivo contido no computador local para dentro do container
//   Entrada:
//     destinationPath: caminho do arquivo no container
//     content: Reader para o arquivo contido no computador. Ex.: err, content = os.Open("/home/user/file.txt")
//   Saída:
//     err: Objeto de erro padrão do Go
//
//   Exemplo:
//     var content io.Reader
//     var destinationPath = "/bin/bash/app.sh"
//     content, err := os.Open("/home/user/app.sh")
//     if err != nil {
//         panic(err)
//     }
//     err = container.ContainerCopyTo(destinationPath, content)
//     if err != nil {
//         panic(err)
//     }
//
// English: Copy a file contained on the local computer into the container
//   Input:
//     destinationPath: file path in container
//     content: Reader for the file contained on the computer. Eg: err, content = os.Open("/home/user/file.txt")
//   Output:
//     err: standard error object
//
//   Example:
//     var content io.Reader
//     var destinationPath = "/bin/bash/app.sh"
//     content, err := os.Open("/home/user/app.sh")
//     if err != nil {
//         panic(err)
//     }
//     err = container.ContainerCopyTo(destinationPath, content)
//     if err != nil {
//         panic(err)
//     }
func (e *ContainerBuilder) ContainerCopyTo(
	destinationPath string,
	content io.Reader,
) (
	err error,
) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.dockerSys.ContainerCopyTo(e.containerID, destinationPath, content)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
