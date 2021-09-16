package iotmakerdockerbuilder

import (
	"time"
)

func (e *ContainerBuilder) StartMonitor(duration *time.Ticker) {
	if e.chaos.monitorStop == nil {
		e.chaos.monitorStop = make(chan struct{}, 1)
	}

	go func() {
		for {
			select {
			case <-e.chaos.monitorStop:
				duration.Stop()
				return

			case <-duration.C:
				e.managerChaos()
			}
		}
	}()
}
