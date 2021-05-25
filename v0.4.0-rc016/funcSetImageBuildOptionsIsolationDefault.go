package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types/container"
)

// SetImageBuildOptionsIsolationDefault (english): Set default isolation mode on current daemon
//
// SetImageBuildOptionsIsolationDefault (português):
func (e *ContainerBuilder) SetImageBuildOptionsIsolationDefault() {
	e.buildOptions.Isolation = container.IsolationDefault
}
