package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) SetSceneName(name string) {
	e.chaos.sceneName = name
}

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

func (e *ContainerBuilder) StopMonitor() {
	if e.chaos.monitorStop == nil {
		return
	}

	e.chaos.monitorStop <- struct{}{}
}
