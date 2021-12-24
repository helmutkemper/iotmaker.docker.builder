package iotmakerdockerbuilder

// EnableChaosScene
//
// English:
//
//  Enables chaos functionality in containers.
//
//   Input:
//     enable: enable chaos manager
//
// Note:
//
//   *This function is used in conjunction with the SetRestartProbability(), SetTimeToStartChaosOnChaosScene(),
//   SetTimeBeforeStartChaosInThisContainerOnChaosScene(), SetTimeOnContainerPausedStateOnChaosScene(), SetTimeOnContainerUnpausedStateOnChaosScene(), SetTimeToRestartThisContainerAfterStopEventOnChaosScene(), StartMonitor() and
//   StopMonitor() functions.
//
// Português:
//
//  Habilita a funcionalidade de caos nos containers.
//
//   Entrada:
//     enable: habilita o gerenciador de caos
//
// Nota:
//
//   * Esta função é usada em conjunto com as funções SetRestartProbability(), SetTimeToStartChaosOnChaosScene(),
//   SetTimeBeforeStartChaosInThisContainerOnChaosScene(), SetTimeOnContainerPausedStateOnChaosScene(), SetTimeOnContainerUnpausedStateOnChaosScene(), SetTimeToRestartThisContainerAfterStopEventOnChaosScene(), StartMonitor() e
//   StopMonitor()
func (e *ContainerBuilder) EnableChaosScene(enable bool) {
	e.chaos.enableChaos = enable
}
