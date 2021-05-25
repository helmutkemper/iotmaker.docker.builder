package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// GetLastInspect
//
// English: Returns the container data based on the last ticker cycle defined in SetInspectInterval()
//
//   Note: the GetChannelOnContainerInspect() function returns the channel triggered by the ticker when the
//   information is ready for use
//
// Português: Retorna os dados do container baseado no último ciclo do ticker definido em SetInspectInterval()
//
//   Nota: a função GetChannelOnContainerInspect() retorna o canal disparado pelo ticker quando as informações estão
//   prontas para uso
//
func (e *ContainerBuilder) GetLastInspect() (inspect iotmakerdocker.ContainerInspect) {
	return e.inspect
}
