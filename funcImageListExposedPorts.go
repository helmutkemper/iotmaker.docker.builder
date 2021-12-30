package iotmakerdockerbuilder

import (
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/util"
)

// ImageListExposedPorts
//
// English:
//
//  Lists all the ports defined in the image to be exposed.
//
//   Output:
//     portList: List of ports exposed on image creation. (Dockerfile expose port)
//     err: standard error object
//
// Note:
//
//   * The ports exposed in the creation of the container can be defined by
//     SetOpenAllContainersPorts(), AddPortToChange() and AddPortToExpose();
//   * By default, all doors are closed.
//
// Português:
//
//  Lista todas as portas definidas na imagem para serem expostas.
//
//   Saída:
//     portList: Lista de portas expostas na criação da imagem. (Dockerfile expose port)
//     err: Objeto de erro padrão
//
// Nota:
//
//   * As portas expostas na criação do container podem ser definidas por SetOpenAllContainersPorts(),
//     AddPortToChange() e AddPortToExpose();
//   * Por padrão, todas as portas ficam fechadas.
func (e *ContainerBuilder) ImageListExposedPorts() (portList []nat.Port, err error) {

	portList, err = e.dockerSys.ImageListExposedPortsByName(e.imageName)
	if err != nil {
		util.TraceToLog()
	}
	return
}
