package iotmakerdockerbuilder

// SetImageBuildOptionsCPUPeriod (english): Specify the CPU CFS scheduler period, which is used alongside --cpu-quota.
// Defaults to 100000 microseconds (100 milliseconds). Most users do not change this from the default.
// For most use-cases, --cpus is a more convenient alternative.
//
// SetImageBuildOptionsCPUPeriod (português):
func (e *ContainerBuilder) SetImageBuildOptionsCPUPeriod(value int64) {
	e.buildOptions.CPUPeriod = value
}
