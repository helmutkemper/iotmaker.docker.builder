package iotmakerdockerbuilder

func (e *ContainerBuilder) AddRestartMatchFlag(value string) {
	if e.chaos.filterRestart == nil {
		e.chaos.filterRestart = make([]LogFilter, 0)
	}

	e.chaos.filterRestart = append(e.chaos.filterRestart, LogFilter{Match: value})
}
