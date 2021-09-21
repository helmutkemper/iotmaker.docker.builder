package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeAggregatePreCPUTimeTheContainerWasThrottled(file *os.File, stats *types.Stats) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if e.rowsToPrint&KAggregatePreCPUTimeTheContainerWasThrottled == KAggregatePreCPUTimeTheContainerWasThrottled {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.ThrottledTime)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KAggregatePreCPUTimeTheContainerWasThrottledComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelAggregatePreCPUTimeTheContainerWasThrottled(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if e.rowsToPrint&KAggregatePreCPUTimeTheContainerWasThrottled == KAggregatePreCPUTimeTheContainerWasThrottled {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Aggregate time the container was throttled for in nanoseconds."))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KAggregatePreCPUTimeTheContainerWasThrottledComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstAggregatePreCPUTimeTheContainerWasThrottled(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if e.rowsToPrint&KAggregatePreCPUTimeTheContainerWasThrottled == KAggregatePreCPUTimeTheContainerWasThrottled {
		_, err = file.Write([]byte("KAggregatePreCPUTimeTheContainerWasThrottled"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KAggregatePreCPUTimeTheContainerWasThrottledComa != 0
	}

	return
}
