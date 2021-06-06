package iotmakerdockerbuilder

// SetImageBuildOptionsCPUSetMems
//
// English: Define a memory nodes (MEMs) (--cpuset-mems)
//
// --cpuset-mems="" Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.
//
// If you have four memory nodes on your system (0-3), use --cpuset-mems=0,1 then processes in your Docker container
// will only use memory from the first two memory nodes.
//
// Português:
func (e *ContainerBuilder) SetImageBuildOptionsCPUSetMems(value string) {
	e.buildOptions.CPUSetMems = value
}
