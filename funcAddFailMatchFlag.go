package iotmakerdockerbuilder

func (e *ContainerBuilder) AddFailMatchFlag(value string) {
	if e.chaos.filterFail == nil {
		e.chaos.filterFail = make([]LogFilter, 0)
	}

	e.chaos.filterFail = append(e.chaos.filterFail, LogFilter{Match: value})
}
