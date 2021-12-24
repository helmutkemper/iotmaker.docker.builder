package iotmakerdockerbuilder

import (
	"time"
)

// StartMonitor
//
// Português: Habilitar um time.Ticker com a finalidade de colher informações de desempenho do container na forma de um
// log CSV e gerencia o controle de caos, caso o mesmo tenha sido habilitado.
//
//   Nota: - Esta função é usada em conjunto com as funções EnableChaosScene() e SetCsvLogPath()
//
// English: Enable a time.Ticker in order to gather performance information from the container in the form of a CSV log
// and manage chaos control, if it has been enabled.
//
//   Nota: - Esta função é usada em conjunto com as funções EnableChaosScene() e SetCsvLogPath()
func (e *ContainerBuilder) StartMonitor() {

	var duration = time.NewTicker(2 * time.Second)

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
