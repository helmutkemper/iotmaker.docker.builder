package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) StartMonitor(duration *time.Ticker) {
	if e.chaos.monitorStop == nil {
		e.chaos.monitorStop = make(chan struct{})
	}

	go func() {
		for {
			select {
			case <-e.chaos.monitorStop:
				return

			case <-duration.C:
				e.managerChaos()
			}
		}
	}()
}
