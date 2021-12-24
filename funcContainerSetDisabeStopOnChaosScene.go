package iotmakerdockerbuilder

// ContainerSetDisabeStopOnChaosScene
//
// English:
//
//  Set the container stop functionality to be disabled when the chaos scene is running
//
//   Entrada:
//     value: true to disable the container stop functionality
//
// Português:
//
//  Define se a funcionalidade de parar o container será desabilitada quando a cena de chaos estiver em execução
//
//   Entrada:
//     value: true para desabilitar a funcionalidade de parar o container
func (e *ContainerBuilder) ContainerSetDisabeStopOnChaosScene(value bool) {
	e.chaos.disableStopContainer = value
}
