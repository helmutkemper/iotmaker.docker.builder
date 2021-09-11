package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeTotalCPUTimeConsumed(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Total CPU time consumed.
	// Units: nanoseconds (Linux)
	// Units: 100's of nanoseconds (Windows)
	if makeLabel == true && e.logFlags&KTotalCPUTimeConsumed == KTotalCPUTimeConsumed {
		_, err = file.Write([]byte("Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTotalCPUTimeConsumed == KTotalCPUTimeConsumed {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.CPUUsage.TotalUsage)))
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
