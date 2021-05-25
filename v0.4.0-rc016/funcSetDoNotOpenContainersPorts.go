package iotmakerdockerbuilder

// SetOpenAllContainersPorts (english):
//
// SetOpenAllContainersPorts (português): Expõe automaticamente todas as portas listadas na imagem usada para gerar o
// container
//
//     Nota: As portas expostas na criação do container pode ser definidas por SetOpenAllContainersPorts(),
//     AddPortToChange() e AddPortToOpen();
//     Por padrão, todas as portas ficam fechadas;
//     A função ImageListExposedPorts() retorna todas as portas definidas na imagem para serem expostas.
//
func (e *ContainerBuilder) SetOpenAllContainersPorts() {
	e.openAllPorts = true
}
