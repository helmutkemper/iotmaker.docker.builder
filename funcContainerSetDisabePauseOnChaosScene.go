package iotmakerdockerbuilder

// ContainerSetDisabePauseOnChaosScene
//
// English:
//
//  Set the container pause functionality to be disabled when the chaos scene is running
//
//   Entrada:
//     value: true to disable the container pause functionality
//
// Português:
//
//  Define se a funcionalidade de pausar o container será desabilitada quando a cena de chaos estiver em execução
//
//   Entrada:
//     value: true para desabilitar a funcionalidade de pausar o container
func (e *ContainerBuilder) ContainerSetDisabePauseOnChaosScene(value bool) {
	e.chaos.disablePauseContainer = value
}
