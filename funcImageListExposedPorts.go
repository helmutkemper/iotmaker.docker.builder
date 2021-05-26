package iotmakerdockerbuilder

import (
	"github.com/docker/go-connections/nat"
)

// ImageListExposedPorts
//
// English: Lists all the ports defined in the image to be exposed.
//
//     Note: The ports exposed in the creation of the container can be defined by SetOpenAllContainersPorts(),
//     AddPortToChange() and AddPortToOpen();
//     By default, all doors are closed.
//
// Português: Lista todas as portas definidas na imagem para serem expostas.
//
//     Nota: As portas expostas na criação do container podem ser definidas por SetOpenAllContainersPorts(),
//     AddPortToChange() e AddPortToOpen();
//     Por padrão, todas as portas ficam fechadas.
func (e *ContainerBuilder) ImageListExposedPorts() (portList []nat.Port, err error) {

	portList, err = e.dockerSys.ImageListExposedPortsByName(e.imageName)
	return
}
