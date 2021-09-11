package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods when the container hits its throttling limit.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.ThrottledPeriods)))
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
