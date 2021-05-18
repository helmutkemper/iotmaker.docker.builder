package iotmaker_docker_builder

// AddPortToOpen (english):
//
// AddPortToOpen (português): Define as portas a serem expostas na rede
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
