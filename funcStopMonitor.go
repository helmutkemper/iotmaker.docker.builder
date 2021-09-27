package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
)

func (e *ContainerBuilder) StopMonitor() (err error) {

	e.chaos.monitorRunning = false

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

		theater.SetContainerUnPaused(e.chaos.sceneName)
		log.Printf("%v: unpause()", e.containerName)
		e.chaos.containerPaused = false

		err = e.ContainerUnpause()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	if e.chaos.containerStopped == true {

		theater.SetContainerUnStopped(e.chaos.sceneName)
		log.Printf("%v: start()", e.containerName)
		e.chaos.containerStopped = false

		err = e.ContainerStart()
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
