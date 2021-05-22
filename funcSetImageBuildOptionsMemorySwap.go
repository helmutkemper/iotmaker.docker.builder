package iotmakerdockerbuilder

// SetImageBuildOptionsMemorySwap (english): Set memory swar (--memory-swap)
//
//   Use value * 1024 = k, value * 1024 * 1024 = m and value * 1024 * 1024 * 1024 = g
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
// SetImageBuildOptionsMemorySwap (português):
//
//   Use value * 1024 = k, value * 1024 * 1024 = m and value * 1024 * 1024 * 1024 = g
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
func (e *ContainerBuilder) SetImageBuildOptionsMemorySwap(value int64) {
	e.buildOptions.MemorySwap = value
}
