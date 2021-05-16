package iotmaker_docker_builder

import (
	"github.com/docker/docker/api/types/mount"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"path/filepath"
)

// AddFiileOrFolderToLinkBetweenConputerHostAndContainer (english):
//
// AddFiileOrFolderToLinkBetweenConputerHostAndContainer (português): Monta um arquivo ou pasta entre o computador e o
// container.
//   computerHostPath:    Caminho do arquivo ou pasta dentro do computador hospedeiro
//   insideContainerPath: Caminho dentro do container
func (e *ContainerBuilder) AddFiileOrFolderToLinkBetweenConputerHostAndContainer(computerHostPath, insideContainerPath string) (err error) {

	if e.volumes == nil {
		e.volumes = make([]mount.Mount, 0)
	}

	computerHostPath, err = filepath.Abs(computerHostPath)
	if err != nil {
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
