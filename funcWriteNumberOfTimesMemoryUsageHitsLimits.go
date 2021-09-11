package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfTimesMemoryUsageHitsLimits(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// number of times memory usage hits limits.
	if makeLabel == true && e.logFlags&KNumberOfTimesMemoryUsageHitsLimits == KNumberOfTimesMemoryUsageHitsLimits {
		_, err = file.Write([]byte("Number of times memory usage hits limits.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfTimesMemoryUsageHitsLimits == KNumberOfTimesMemoryUsageHitsLimits {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Failcnt)))
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
