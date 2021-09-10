package iotmakerdockerbuilder

func (e *ContainerBuilder) AddFilterToRestartContainer(match, filter, search, replace string) {
	if e.chaos.filterRestart == nil {
		e.chaos.filterRestart = make([]LogFilter, 0)
	}

	e.chaos.filterRestart = append(
		e.chaos.filterRestart,
		LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)
}
