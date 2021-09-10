package iotmakerdockerbuilder

func (e *ContainerBuilder) AddFilterToLog(label, match, filter, search, replace string) {
	if e.chaos.filterLog == nil {
		e.chaos.filterLog = make([]LogFilter, 0)
	}

	e.chaos.filterLog = append(
		e.chaos.filterLog,
		LogFilter{
			Label:   label,
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)
}
