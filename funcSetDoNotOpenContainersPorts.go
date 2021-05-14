package iotmaker_docker_builder

// SetDoNotOpenContainersPorts (english):
//
// SetDoNotOpenContainersPorts (português): Impede a publicação de portas expostas na rede de forma automática
//
//     Nota: A omissão de definição das portas a serem expostas define automaticamente todas as portas contidas na
//     imagem como portas a serem expostas.
//     AddPortToOpen() e AddPortToChange() limitam as portas abertas as portas listadas.
//     SetDoNotOpenContainersPorts() impede a exposição automática de portas.
func (e *ContainerBuilder) SetDoNotOpenContainersPorts() {
	e.doNotOpenPorts = true
}
