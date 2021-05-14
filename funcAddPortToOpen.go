package iotmaker_docker_builder

// AddPortToOpen (english):
//
// AddPortToOpen (português): Define as portas a serem expostas na rede
//   value: porta na forma de string numérica
//
//     Nota: A omissão de definição das portas a serem expostas define automaticamente todas as portas contidas na
//     imagem como portas a serem expostas.
//     AddPortToOpen() e AddPortToChange() limitam as portas abertas as portas listadas.
//     SetDoNotOpenContainersPorts() impede a exposição automática de portas.
func (e *ContainerBuilder) AddPortToOpen(value string) {
	if e.openPorts == nil {
		e.openPorts = make([]string, 0)
	}

	e.openPorts = append(e.openPorts, value)
}
