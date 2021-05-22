package iotmakerdockerbuilder

// SetImageBuildOptionsCPUQuota (english): Set this flag to a value greater or less than the default of 1024 to increase
// or reduce the container’s weight, and give it access to a greater or lesser proportion of the host machine’s CPU
// cycles.
//
// This is only enforced when CPU cycles are constrained. When plenty of CPU cycles are available, all containers use as
// much CPU as they need. In that way, this is a soft limit. --cpu-shares does not prevent containers from being
// scheduled in swarm mode. It prioritizes container CPU resources for the available CPU cycles. It does not guarantee
// or reserve any specific CPU access.
//
// SetImageBuildOptionsCPUQuota (português):
func (e *ContainerBuilder) SetImageBuildOptionsCPUQuota(value int64) {
	e.buildOptions.CPUQuota = value
}
