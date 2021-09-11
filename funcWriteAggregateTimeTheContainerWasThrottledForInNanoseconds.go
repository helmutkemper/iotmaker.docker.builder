package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeAggregateTimeTheContainerWasThrottledForInNanoseconds(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if makeLabel == true && e.logFlags&KAggregateTimeTheContainerWasThrottledForInNanoseconds == KAggregateTimeTheContainerWasThrottledForInNanoseconds {
		_, err = file.Write([]byte("Throttling Data. Linux only. Aggregate time the container was throttled for in nanoseconds.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KAggregateTimeTheContainerWasThrottledForInNanoseconds == KAggregateTimeTheContainerWasThrottledForInNanoseconds {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.ThrottledTime)))
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
