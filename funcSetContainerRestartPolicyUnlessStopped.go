package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

// SetContainerRestartPolicyUnlessStopped (english):
//
// SetContainerRestartPolicyUnlessStopped (português): Define a política de reinício do container como sempre reinicia o container, caso ele não tenha sido parado manualmente.
func (e *ContainerBuilder) SetContainerRestartPolicyUnlessStopped() {
	e.restartPolicy = iotmakerdocker.KRestartPolicyUnlessStopped
}
