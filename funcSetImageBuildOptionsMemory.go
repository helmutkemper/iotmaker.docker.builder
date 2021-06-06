package iotmakerdockerbuilder

// SetImageBuildOptionsMemory
//
// English: The maximum amount of memory the container can use.
// If you set this option, the minimum allowed value is 4 * 1024 * 1024 (4 megabyte).
//
//   Use value * KKiloByte, value * KMegaByte and value * KGigaByte
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
//
// Português: Memória máxima total que o container pode usar.
// Se você vai usar esta opção, o máximo permitido é 4 * 1024 * 1024 (4 megabyte)
//
//   Use value * KKiloByte, value * KMegaByte e value * KGigaByte
//   See https://docs.docker.com/engine/reference/run/#user-memory-constraints
func (e *ContainerBuilder) SetImageBuildOptionsMemory(value int64) {
	e.buildOptions.Memory = value
}
