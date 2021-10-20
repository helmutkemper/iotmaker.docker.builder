package iotmakerdockerbuilder

// AddRestartMatchFlag
//
//
//
// Português: Adiciona um texto ao filtro de reinício do container.
//
// Durante o teste de caos, este filtro procura um texto na saída padrão do container e quando este texto é encontrado,
// o indicador de reinício libera o container para ser reiniciado durante o teste, porém, a escolha de reiniciar o
// container ou não será feita aleatóriamente durante o teste de caos, e só nele.
func (e *ContainerBuilder) AddRestartMatchFlag(value string) {
	if e.chaos.filterRestart == nil {
		e.chaos.filterRestart = make([]LogFilter, 0)
	}

	e.chaos.filterRestart = append(e.chaos.filterRestart, LogFilter{Match: value})
}
