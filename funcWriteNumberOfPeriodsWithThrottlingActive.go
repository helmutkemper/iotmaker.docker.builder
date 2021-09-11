package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfPeriodsWithThrottlingActive(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWithThrottlingActive == KNumberOfPeriodsWithThrottlingActive {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods with throttling active.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWithThrottlingActive == KNumberOfPeriodsWithThrottlingActive {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.Periods)))
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
