package iotmakerdockerbuilder

// AddPortToOpen
//
// English: Defines the port to be exposed on the network
//
//   value: port in the form of a numeric string
//
//     Note: The ports exposed in the creation of the container can be defined by SetOpenAllContainersPorts(),
//     AddPortToChange() and AddPortToOpen();
//     By default, all doors are closed;
//
//     The ImageListExposedPorts() function returns all ports defined in the image to be exposed.
//
// Português: Define a porta a ser expostas na rede
//
//   value: porta na forma de string numérica
//
//     Nota: As portas expostas na criação do container pode ser definidas por SetOpenAllContainersPorts(),
//     AddPortToChange() e AddPortToOpen();
//     Por padrão, todas as portas ficam fechadas;
//     A função ImageListExposedPorts() retorna todas as portas definidas na imagem para serem expostas.
//
func (e *ContainerBuilder) AddPortToOpen(value string) {
	if e.openPorts == nil {
		e.openPorts = make([]string, 0)
	}

	e.openPorts = append(e.openPorts, value)
}
