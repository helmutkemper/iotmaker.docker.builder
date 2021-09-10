package iotmakerdockerbuilder

func (e *ContainerBuilder) AddFilterToStartChaos(match, filter, search, replace string) {
	if e.chaos.filterToStart == nil {
		e.chaos.filterToStart = make([]LogFilter, 0)
	}

	e.chaos.filterToStart = append(
		e.chaos.filterToStart,
		LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)
}
