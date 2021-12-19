package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"io"
	"io/ioutil"
)

// ContainerCopyFrom
//
// Português: Copia um arquivo contido no container para uma pasta local
//   Entrada:
//     containerPathList: lista de arquivos contidos no container
//     hostPathList:      lista de caminhos dos arquivos a serem salvos no host
//   Saída:
//     statsList: Lista de informações dos arquivos
//     err:       Objeto padrão de error
//
// English: Copy a file contained in the container to a local folder
//   Input:
//     containerPathList: list of files contained in the container
//     hostPathList:      list of file paths to be saved on the host
//   Output:
//     statsList: List of file information
//     err:       Default error object
func (e *ContainerBuilder) ContainerCopyFrom(
	containerPathList []string,
	hostPathList []string,
) (
	statsList []types.ContainerPathStat,
	err error,
) {
	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	var reader io.ReadCloser
	var stats types.ContainerPathStat
	var data []byte
	for k, sourcePath := range containerPathList {
		reader, stats, err = e.dockerSys.ContainerCopyFrom(e.containerID, sourcePath)
		if err != nil {
			util.TraceToLog()
			return
		}
		data, err = ioutil.ReadAll(reader)
		if err != nil {
			util.TraceToLog()
			return
		}
		err = ioutil.WriteFile(hostPathList[k], data, 0777)
		if err != nil {
			util.TraceToLog()
			return
		}

		statsList = append(statsList, stats)
	}

	return
}
