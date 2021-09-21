package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeAggregateTimeTheContainerWasThrottledForInNanoseconds(file *os.File, stats *types.Stats) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if e.rowsToPrint&KAggregateTimeTheContainerWasThrottledForInNanoseconds == KAggregateTimeTheContainerWasThrottledForInNanoseconds {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.ThrottledTime)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KAggregateTimeTheContainerWasThrottledForInNanosecondsComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelAggregateTimeTheContainerWasThrottledForInNanoseconds(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if e.rowsToPrint&KAggregateTimeTheContainerWasThrottledForInNanoseconds == KAggregateTimeTheContainerWasThrottledForInNanoseconds {
		_, err = file.Write([]byte("Throttling Data. Linux only. Aggregate time the container was throttled for in nanoseconds."))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KAggregateTimeTheContainerWasThrottledForInNanosecondsComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstAggregateTimeTheContainerWasThrottledForInNanoseconds(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if e.rowsToPrint&KAggregateTimeTheContainerWasThrottledForInNanoseconds == KAggregateTimeTheContainerWasThrottledForInNanoseconds {
		_, err = file.Write([]byte("KAggregateTimeTheContainerWasThrottledForInNanoseconds"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KAggregateTimeTheContainerWasThrottledForInNanosecondsComa != 0
	}

	return
}
