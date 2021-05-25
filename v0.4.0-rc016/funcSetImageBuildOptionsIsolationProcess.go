package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types/container"
)

// SetImageBuildOptionsIsolationProcess (english): Set process isolation mode
//
// SetImageBuildOptionsIsolationProcess (português):
func (e *ContainerBuilder) SetImageBuildOptionsIsolationProcess() {
	e.buildOptions.Isolation = container.IsolationProcess
}
