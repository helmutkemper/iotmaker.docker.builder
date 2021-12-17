package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"io"
)

// ContainerCopyFrom
//
// Português: Oopia um arquivo contido no container para uma pasta local
//   Entrada:
//     sourcePath: caminho do arquivo no container
//   Saída:
//     reader: Reader para o arquivo contido no container. Ex.: fileInBytesArr, err = ioutil.ReadAll(reader)
//     stats:  Informações do arquivo
//     err:    Objeto padrão de error do Go
//
//   Exemplo:
//      var reader io.Reader
//	    var data []byte
//	    reader, _, err = container.ContainerCopyFrom("/go/bin/golangci-lint")
//	    if err != nil {
//		    panic(err)
//	    }
//	    data, err = ioutil.ReadAll(reader)
//	    if err != nil {
//		    panic(err)
//	    }
//      err = ioutil.WriteFile("./alpineFiles/golangci-lint", data, 0777)
//	    if err != nil {
//		    panic(err)
//	    }
//
// English: Copy a file contained in the container to a local folder
//   Input:
//     sourcePath: file path in container
//   Output:
//     reader: Reader for the file contained in the container. Eg: fileInBytesArr, err = ioutil.ReadAll(reader)
//     stats:  file information
//     err:    Go's default error object
//
//   Example
//      var reader io.Reader
//	    var data []byte
//	    reader, _, err = container.ContainerCopyFrom("/go/bin/golangci-lint")
//	    if err != nil {
//		    panic(err)
//	    }
//	    data, err = ioutil.ReadAll(reader)
//	    if err != nil {
//		    panic(err)
//	    }
//      err = ioutil.WriteFile("./alpineFiles/golangci-lint", data, 0777)
//	    if err != nil {
//		    panic(err)
//	    }
func (e *ContainerBuilder) ContainerCopyFrom(
	sourcePath string,
) (
	reader io.ReadCloser,
	stats types.ContainerPathStat,
	err error,
) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	reader, stats, err = e.dockerSys.ContainerCopyFrom(e.containerID, sourcePath)
	return
}
