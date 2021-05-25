package iotmakerdockerbuilder

// SetImageBuildOptionsCPUSetCPUs (english): Limit the specific CPUs or cores a container can use. A comma-separated
// list or hyphen-separated range of CPUs a container can use, if you have more than one CPU. The first CPU is
// numbered 0. A valid value might be 0-3 (to use the first, second, third, and fourth CPU) or 1,3
// (to use the second and fourth CPU).
//
// SetImageBuildOptionsCPUSetCPUs (português):
func (e *ContainerBuilder) SetImageBuildOptionsCPUSetCPUs(value string) {
	e.buildOptions.CPUSetCPUs = value
}
