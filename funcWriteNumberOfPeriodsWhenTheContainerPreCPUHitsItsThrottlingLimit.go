package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Number of periods when the container hits its throttling limit.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.ThrottledPeriods)))
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
