package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// SetContainerRestartPolicyNo (english):
//
// SetContainerRestartPolicyNo (português): Define a política de reinício do container como não reiniciar o container (padrão).
func (e *ContainerBuilder) SetContainerRestartPolicyNo() {
	e.restartPolicy = iotmakerdocker.KRestartPolicyNo
}
