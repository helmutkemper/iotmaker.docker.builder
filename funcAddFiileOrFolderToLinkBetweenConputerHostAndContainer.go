package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types/mount"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
	"path/filepath"
)

// AddFileOrFolderToLinkBetweenConputerHostAndContainer
//
// English: Links a file or folder between the computer host and the container.
//
//   computerHostPath:    Path of the file or folder inside the host computer
//   insideContainerPath: Path inside the container
//
// Português: Vincula um arquivo ou pasta entre o computador e o container.
//
//   computerHostPath:    Caminho do arquivo ou pasta no computador hospedeiro
//   insideContainerPath: Caminho dentro do container
func (e *ContainerBuilder) AddFileOrFolderToLinkBetweenConputerHostAndContainer(computerHostPath, insideContainerPath string) (err error) {

	if e.volumes == nil {
		e.volumes = make([]mount.Mount, 0)
	}

	computerHostPath, err = filepath.Abs(computerHostPath)
	if err != nil {
		util.TraceToLog()
		return
	}

	e.volumes = append(
		e.volumes,
		mount.Mount{
			// bind - is the type for mounting host dir (real folder inside computer where this code work)
			Type: iotmakerdocker.KVolumeMountTypeBindString,
			// path inside host machine
			Source: computerHostPath,
			// path inside image
			Target: insideContainerPath,
		},
	)

	return
}
