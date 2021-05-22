package iotmakerdockerbuilder

// SetImageBuildOptionsMemory (english): The maximum amount of memory the container can use. If you set this option, the minimum allowed value is 4 * 1024 * 1024 (4 megabyte).
//
//   Use value * 1024 = k, value * 1024 * 1024 = m and value * 1024 * 1024 * 1024 = g
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
// SetImageBuildOptionsMemory (português):
//
//   Use value * 1024 = k, value * 1024 * 1024 = m and value * 1024 * 1024 * 1024 = g
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
func (e *ContainerBuilder) SetImageBuildOptionsMemory(value int64) {
	e.buildOptions.Memory = value
}
