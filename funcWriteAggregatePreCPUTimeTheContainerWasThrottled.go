package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeAggregatePreCPUTimeTheContainerWasThrottled(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if makeLabel == true && e.logFlags&KAggregatePreCPUTimeTheContainerWasThrottled == KAggregatePreCPUTimeTheContainerWasThrottled {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Aggregate time the container was throttled for in nanoseconds.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KAggregatePreCPUTimeTheContainerWasThrottled == KAggregatePreCPUTimeTheContainerWasThrottled {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.ThrottledTime)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	return
}
