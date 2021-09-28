package iotmakerdockerbuilder

import (
	"time"
)

func (e *ContainerBuilder) StartMonitor(duration *time.Ticker) {

	if e.chaos.monitorRunning == true {
		return
	}

	e.chaos.monitorRunning = true

	if e.chaos.monitorStop == nil {
		e.chaos.monitorStop = make(chan struct{}, 1)
	}

	go func() {
		for {
			select {
			case <-e.chaos.monitorStop:
				duration.Stop()
				_ = e.stopMonitorAfterStopped()
				return

			case <-duration.C:
				e.managerChaos()

				if e.chaos.monitorRunning == false {
					duration.Stop()
					_ = e.stopMonitorAfterStopped()
					return
				}
			}
		}
	}()
}
