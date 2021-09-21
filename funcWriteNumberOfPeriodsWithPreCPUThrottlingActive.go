package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfPeriodsWithPreCPUThrottlingActive(file *os.File, stats *types.Stats) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if e.rowsToPrint&KNumberOfPeriodsWithPreCPUThrottlingActive == KNumberOfPeriodsWithPreCPUThrottlingActive {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.Periods)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWithPreCPUThrottlingActiveComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelNumberOfPeriodsWithPreCPUThrottlingActive(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if e.rowsToPrint&KNumberOfPeriodsWithPreCPUThrottlingActive == KNumberOfPeriodsWithPreCPUThrottlingActive {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Number of periods with throttling active."))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWithPreCPUThrottlingActiveComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstNumberOfPeriodsWithPreCPUThrottlingActive(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if e.rowsToPrint&KNumberOfPeriodsWithPreCPUThrottlingActive == KNumberOfPeriodsWithPreCPUThrottlingActive {
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
