package iotmakerdockerbuilder

import "time"

// SetTimeToStartChaosOnChaosScene
//
// English:
//
//  This function sets a timeout before the chaos test starts, when indicator text is encountered in the standard output.
//
//   Input:
//     min: minimum waiting time until chaos test starts
//     max: maximum waiting time until chaos test starts
//
// Basically, the idea is that you put at some point in the test a text like, chaos can be initialized, in the container's standard output and the time gives a random character to when the chaos starts.
//
// Note:
//
//   * This function is used in conjunction with the AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() or AddFilterToStartChaos() functions
//
// Português:
//
//  Esta função define um tempo de espera antes do teste de caos começar, quando o texto indicador é incontrado na saída padrão.
//
//   Entrada:
//     min: tempo mínimo de espera até o teste de caos começar
//     max: tempo máximo de espera até o teste de caos começar
//
// Basicamente, a ideia é que você coloque em algum ponto do teste um texto tipo, caos pode ser inicializado, na saída padrão do container e o tempo dá um caráter aleatório a quando o caos começa.
//
// Nota:
//
//   * Esta função é usada em conjunto com as funções AddStartChaosMatchFlag(), AddStartChaosMatchFlagToFileLog() ou AddFilterToStartChaos()
func (e *ContainerBuilder) SetTimeToStartChaosOnChaosScene(min, max time.Duration) {
	e.chaos.minimumTimeToStartChaos = min
	e.chaos.maximumTimeToStartChaos = max
}
