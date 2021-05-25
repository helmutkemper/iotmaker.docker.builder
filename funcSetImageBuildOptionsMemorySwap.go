package iotmakerdockerbuilder

// SetImageBuildOptionsMemorySwap (english): Set memory swar (--memory-swap)
//
//   Use value * KKiloByte, value * KMegaByte and value * KGigaByte
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
// SetImageBuildOptionsMemorySwap (português):
//
//   Use value * KKiloByte, value * KMegaByte and value * KGigaByte
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
func (e *ContainerBuilder) SetImageBuildOptionsMemorySwap(value int64) {
	e.buildOptions.MemorySwap = value
}
