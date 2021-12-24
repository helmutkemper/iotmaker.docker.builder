package iotmakerdockerbuilder

import "time"

// SetTimeOnContainerPausedStateOnChaosScene
//
// English:
//
//  Sets the minimum and maximum times for the container pause
//
//   Input:
//     min: minimum time for container pause
//     max: maximum time for container pause
//
// Note:
//
//   * This function is used in conjunction with the AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() or AddFilterToStartChaos() functions
//
// Português:
//
//  Define os tempos mínimos e máximos para a pausa do container
//
//   Entrada:
//     min: tempo mínimo para a pausa do container
//     max: tempo máximo para a pausa do container
//
// Nota:
//
//   * Esta função é usada em conjunto com as funções AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() ou AddFilterToStartChaos()
func (e *ContainerBuilder) SetTimeOnContainerPausedStateOnChaosScene(min, max time.Duration) {
	e.chaos.minimumTimeToPause = min
	e.chaos.maximumTimeToPause = max
}
