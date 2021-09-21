package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit(file *os.File, stats *types.Stats) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
	if e.rowsToPrint&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.ThrottledPeriods)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimitComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
	if e.rowsToPrint&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods when the container hits its throttling limit."))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimitComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit(file *os.File) (tab bool, err error) {
	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
	if e.rowsToPrint&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit {
		_, err = file.Write([]byte("KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimitComa != 0
	}

	return
}
