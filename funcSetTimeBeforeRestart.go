package iotmakerdockerbuilder

import "time"

// SetTimeBeforeStartChaosInThisContainerOnChaosScene
//
// English:
//
//  Defines the minimum and maximum waiting times before enabling the restart of containers in a chaos scenario
//
//  The choice of time will be made randomly between the minimum and maximum values
//
//   Input:
//     min: minimum waiting time
//     max: maximum wait time
//
// Note:
//
//   * This function is used in conjunction with the AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() or AddFilterToStartChaos() functions
//
// Português:
//
//  Define os tempos mínimo e máximos de espera antes de habilitar o reinício dos containers em um cenário de caos
//
//  A escolha do tempo será feita de forma aleatória entre os valores mínimo e máximo
//
//   Entrada:
//     min: tempo de espera mínimo
//     max: tempo de espera máximo
//
// Nota:
//
//   * Esta função é usada em conjunto com as funções AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() ou AddFilterToStartChaos()
//
func (e *ContainerBuilder) SetTimeBeforeStartChaosInThisContainerOnChaosScene(min, max time.Duration) {
	e.chaos.minimumTimeBeforeRestart = min
	e.chaos.maximumTimeBeforeRestart = max
}
