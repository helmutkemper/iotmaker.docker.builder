package iotmakerdockerbuilder

func (e *ContainerBuilder) AddSuccessMatchFlag(value string) {
	if e.chaos.filterSuccess == nil {
		e.chaos.filterSuccess = make([]LogFilter, 0)
	}

	e.chaos.filterSuccess = append(e.chaos.filterSuccess, LogFilter{Match: value})
}
