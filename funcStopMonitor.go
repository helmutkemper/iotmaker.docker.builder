package iotmakerdockerbuilder

func (e *ContainerBuilder) StopMonitor() (err error) {

	e.chaos.monitorRunning = false

	if len(e.chaos.monitorStop) == 0 {
		e.chaos.monitorStop <- struct{}{}
	}

	return
}
