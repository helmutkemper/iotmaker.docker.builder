package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// SetContainerRestartPolicyNo
//
// English: Do not automatically restart the container. (the default)
//
// Português: Define a política de reinício do container como não reiniciar o container (padrão).
func (e *ContainerBuilder) SetContainerRestartPolicyNo() {
	e.restartPolicy = iotmakerdocker.KRestartPolicyNo
}
