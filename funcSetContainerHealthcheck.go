package iotmakerdockerbuilder

// SetContainerHealthcheck (english):
//
// SetContainerHealthcheck (português):
func (e *ContainerBuilder) SetContainerHealthcheck(value *HealthConfig) {
	e.containerConfig.Healthcheck.Test = value.Test
	e.containerConfig.Healthcheck.Interval = value.Interval
	e.containerConfig.Healthcheck.Timeout = value.Timeout
	e.containerConfig.Healthcheck.StartPeriod = value.StartPeriod
	e.containerConfig.Healthcheck.Retries = value.Retries
}
