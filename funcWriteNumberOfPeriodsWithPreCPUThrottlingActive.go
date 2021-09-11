package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfPeriodsWithPreCPUThrottlingActive(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWithPreCPUThrottlingActive == KNumberOfPeriodsWithPreCPUThrottlingActive {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Number of periods with throttling active.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWithPreCPUThrottlingActive == KNumberOfPeriodsWithPreCPUThrottlingActive {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.Periods)))
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
