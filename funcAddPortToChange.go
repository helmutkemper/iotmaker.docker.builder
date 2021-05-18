package iotmakerdockerbuilder

// AddPortToChange (english):
//
// AddPortToChange (português): Define as portas a serem expostas na rede alterando o valor da porta definida na imagem
// e o valor exposto na rede
//   imagePort: porta definida na imagem, na forma de string numérica
//   newPort: novo valor da porta a se exposta na rede
//
//     Nota: As portas expostas na criação do container pode ser definidas por SetOpenAllContainersPorts(),
//     AddPortToChange() e AddPortToOpen();
//     Por padrão, todas as portas ficam fechadas;
//     A função ImageListExposedPorts() retorna todas as portas definidas na imagem para serem expostas.
//
func (e *ContainerBuilder) AddPortToChange(imagePort string, newPort string) {
	if e.changePorts == nil {
		e.changePorts = make([]changePort, 0)
	}

	e.changePorts = append(
		e.changePorts,
		changePort{
			oldPort: imagePort,
			newPort: newPort,
		},
	)
}
