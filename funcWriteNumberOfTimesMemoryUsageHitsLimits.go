package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfTimesMemoryUsageHitsLimits(file *os.File, stats *types.Stats) (tab bool, err error) {
	// number of times memory usage hits limits.
	if e.rowsToPrint&KNumberOfTimesMemoryUsageHitsLimits == KNumberOfTimesMemoryUsageHitsLimits {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Failcnt)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfTimesMemoryUsageHitsLimitsComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelNumberOfTimesMemoryUsageHitsLimits(file *os.File) (tab bool, err error) {
	// number of times memory usage hits limits.
	if e.rowsToPrint&KNumberOfTimesMemoryUsageHitsLimits == KNumberOfTimesMemoryUsageHitsLimits {
		_, err = file.Write([]byte("Number of times memory usage hits limits."))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfTimesMemoryUsageHitsLimitsComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstNumberOfTimesMemoryUsageHitsLimits(file *os.File) (tab bool, err error) {
	// number of times memory usage hits limits.
	if e.rowsToPrint&KNumberOfTimesMemoryUsageHitsLimits == KNumberOfTimesMemoryUsageHitsLimits {
		_, err = file.Write([]byte("KNumberOfTimesMemoryUsageHitsLimits"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfTimesMemoryUsageHitsLimitsComa != 0
	}

	return
}
