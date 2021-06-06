package iotmakerdockerbuilder

// SetImageBuildOptionsMemorySwap
//
// English: Set memory swap (--memory-swap)
//
//   Use value * KKiloByte, value * KMegaByte and value * KGigaByte
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
// Português: habilita a opção memory swp
//
//   Use value * KKiloByte, value * KMegaByte e value * KGigaByte
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
func (e *ContainerBuilder) SetImageBuildOptionsMemorySwap(value int64) {
	e.buildOptions.MemorySwap = value
}
