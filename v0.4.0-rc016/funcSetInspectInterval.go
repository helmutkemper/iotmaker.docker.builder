package iotmakerdockerbuilder

import (
	"time"
)

// SetInspectInterval (english):
//
// SetInspectInterval (português): Define o intervalo de monitoramento do container [opcional]
//   value: intervalo de tempo entre os eventos de inspeção do container
//
//     Nota: Esta função tem um custo computacional elevado e deve ser usada com moderação.
//     Os valores capturados são apresentados por GetLastInspect() e GetChannelOnContainerInspect()
func (e *ContainerBuilder) SetInspectInterval(value time.Duration) {
	e.inspectInterval = value
}
