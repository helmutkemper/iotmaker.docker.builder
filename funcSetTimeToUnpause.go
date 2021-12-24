package iotmakerdockerbuilder

import "time"

// SetTimeOnContainerUnpausedStateOnChaosScene
//
// English:
//
//  Defines the minimum and maximum times where the container is kept out of the paused state
//
//   Input:
//     min: minimum time out of sleep state
//     max: maximum time out of sleep state
//
// Note:
//
//   * This function is used in conjunction with the AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() or AddFilterToStartChaos() functions
//
// Português:
//
//  Define os tempos mínimos e máximos onde o container é mantido fora do estado de pausa
//
//   Entrada:
//     min: tempo mínimo fora do estado de pausa
//     max: tempo máximo fora do estado de pausa
//
// Nota:
//
//   * Esta função é usada em conjunto com as funções AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() ou AddFilterToStartChaos()
func (e *ContainerBuilder) SetTimeOnContainerUnpausedStateOnChaosScene(min, max time.Duration) {
	e.chaos.minimumTimeToUnpause = min
	e.chaos.maximumTimeToUnpause = max
}
