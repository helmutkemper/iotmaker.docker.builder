package iotmaker_docker_builder

import (
	"github.com/docker/go-connections/nat"
)

// ImageListExposedPorts (english):
//
// ImageListExposedPorts (português): Lista todas as portas definidas na imagem para serem expostas.
//
//     Nota: As portas expostas na criação do container pode ser definidas por SetOpenAllContainersPorts(),
//     AddPortToChange() e AddPortToOpen();
//     Por padrão, todas as portas ficam fechadas;
//     A função ImageListExposedPorts() retorna todas as portas definidas na imagem para serem expostas.
//
func (e *ContainerBuilder) ImageListExposedPorts() (portList []nat.Port, err error) {

	portList, err = e.dockerSys.ImageListExposedPortsByName(e.imageName)
	return
}
