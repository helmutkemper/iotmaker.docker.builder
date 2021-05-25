package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types/container"
)

// SetImageBuildOptionsIsolationHyperV (english): Set HyperV isolation mode
//
// SetImageBuildOptionsIsolationHyperV (português):
func (e *ContainerBuilder) SetImageBuildOptionsIsolationHyperV() {
	e.buildOptions.Isolation = container.IsolationHyperV
}
