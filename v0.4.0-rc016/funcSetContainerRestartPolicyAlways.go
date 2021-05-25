package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// SetContainerRestartPolicyAlways (english):
//
// SetContainerRestartPolicyAlways (português): Define a política de reinício do container como sempre reinicia o container quando ele para, mesmo quando ele é parado manualmente.
func (e *ContainerBuilder) SetContainerRestartPolicyAlways() {
	e.restartPolicy = iotmakerdocker.KRestartPolicyAlways
}
