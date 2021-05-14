package iotmaker_docker_builder

// AddPortToChange (english):
//
// AddPortToChange (português): Define as portas a serem expostas na rede alterando o valor da porta definida na imagem
// e o valor exposto na rede
//   imagePort: porta definida na imagem, na forma de string numérica
//   newPort: novo valor da porta a se exposta na rede
//
//     Nota: A omissão de definição das portas a serem expostas define automaticamente todas as portas contidas na
//     imagem como portas a serem expostas.
//     AddPortToOpen() e AddPortToChange() limitam as portas abertas as portas listadas.
//     SetDoNotOpenContainersPorts() impede a exposição automática de portas.
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
