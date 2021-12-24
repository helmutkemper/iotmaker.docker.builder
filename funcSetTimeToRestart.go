package iotmakerdockerbuilder

import "time"

// SetTimeToRestartThisContainerAfterStopEventOnChaosScene
//
// English
//
//  Defines the minimum and maximum times to restart the container after the container stop event.
//
//   Input:
//     min: minimum timeout before restarting container
//     max: maximum timeout before restarting container
//
// Note:
//
//   * This function is used in conjunction with the AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() or AddFilterToStartChaos() functions
//
// Português:
//
//  Define os tempos mínimos e máximos para reiniciar o container após o evento de parar container.
//
//   Entrada:
//     min: tempo mínimo de espera antes de reiniciar o container
//     max: tempo máximo de espera antes de reiniciar o container
//
// Nota:
//
//   * Esta função é usada em conjunto com as funções AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() ou AddFilterToStartChaos()
func (e *ContainerBuilder) SetTimeToRestartThisContainerAfterStopEventOnChaosScene(min, max time.Duration) {
	e.chaos.minimumTimeToRestart = min
	e.chaos.maximumTimeToRestart = max
}
