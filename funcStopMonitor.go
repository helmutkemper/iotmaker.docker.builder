package iotmakerdockerbuilder

import "github.com/helmutkemper/util"

func (e *ContainerBuilder) StopMonitor() (err error) {
	if e.chaos.monitorStop == nil {
		return
	}

	if e.containerID == "" {
		err = e.GetIdByContainerName()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	e.chaos.linear = true

	if e.chaos.containerPaused == true {
		err = e.ContainerUnpause()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	if e.chaos.containerStopped == true {
		err = e.ContainerRestart()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	if len(e.chaos.monitorStop) == 0 {
		e.chaos.monitorStop <- struct{}{}
	}

	return
}
