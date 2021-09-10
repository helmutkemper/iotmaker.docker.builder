package iotmakerdockerbuilder

func (e *ContainerBuilder) AddFilterToSuccess(match, filter, search, replace string) {
	if e.chaos.filterSuccess == nil {
		e.chaos.filterSuccess = make([]LogFilter, 0)
	}

	e.chaos.filterSuccess = append(
		e.chaos.filterSuccess,
		LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)
}
