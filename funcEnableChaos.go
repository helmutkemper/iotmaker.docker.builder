package iotmakerdockerbuilder

// EnableChaos
//
// English: Enables chaos functionality in containers.
//   Input:
//     enable: enable chaos manager
//
//   Note: - This function is used in conjunction with the SetRestartProbability(), SetTimeToStartChaos(),
//           SetTimeBeforeRestart(), SetTimeToPause(), SetTimeToUnpause(), SetTimeToRestart(), StartMonitor() and
//           StopMonitor() functions.
//
// Português: Habilita a funcionalidade de caos nos containers.
//   Entrada:
//     enable: habilita o gerenciador de caos
//
//   Nota: - Esta função é usada em conjunto com as funções SetRestartProbability(), SetTimeToStartChaos(),
//           SetTimeBeforeRestart(), SetTimeToPause(), SetTimeToUnpause(), SetTimeToRestart(), StartMonitor() e
//           StopMonitor()
func (e *ContainerBuilder) EnableChaos(enable bool) {
	e.chaos.enableChaos = enable
}
