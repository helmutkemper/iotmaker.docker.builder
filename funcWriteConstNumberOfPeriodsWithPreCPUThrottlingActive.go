package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeConstNumberOfPeriodsWithPreCPUThrottlingActive(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if e.rowsToPrint&KLogColumnNumberOfPeriodsWithPreCPUThrottlingActive == KLogColumnNumberOfPeriodsWithPreCPUThrottlingActive {
		_, err = file.Write([]byte("KNumberOfPeriodsWithPreCPUThrottlingActive"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWithPreCPUThrottlingActiveComa != 0
	}

	return
}
