package iotmaker_docker_builder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// GetLastInspect (english):
//
// GetLastInspect (português): Retorna os dados do container baseado no último ciclo do ticker definido em
// SetInspectInterval()
//
//   Nota: a função GetChannelOnContainerInspect() retorna o canal disparado pelo ticker quando as informações estão
//   prontas para uso
//
func (e *ContainerBuilder) GetLastInspect() (inspect iotmakerdocker.ContainerInspect) {
	return e.inspect
}
