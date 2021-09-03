package iotmakerdockerbuilder

type Event struct {
	ContainerName string
	Message       string
	Error         bool
	Done          bool
	Fail          bool
}

func (e *Event) clear() {
	e.ContainerName = ""
	e.Message = ""
	e.Done = false
	e.Error = false
	e.Fail = false
}
