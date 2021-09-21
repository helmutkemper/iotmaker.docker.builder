package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfPeriodsWithThrottlingActive(file *os.File, stats *types.Stats) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if e.rowsToPrint&KNumberOfPeriodsWithThrottlingActive == KNumberOfPeriodsWithThrottlingActive {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.Periods)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWithThrottlingActiveComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelNumberOfPeriodsWithThrottlingActive(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if e.rowsToPrint&KNumberOfPeriodsWithThrottlingActive == KNumberOfPeriodsWithThrottlingActive {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods with throttling active."))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWithThrottlingActiveComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstNumberOfPeriodsWithThrottlingActive(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if e.rowsToPrint&KNumberOfPeriodsWithThrottlingActive == KNumberOfPeriodsWithThrottlingActive {
		_, err = file.Write([]byte("KNumberOfPeriodsWithThrottlingActive"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWithThrottlingActiveComa != 0
	}

	return
}
