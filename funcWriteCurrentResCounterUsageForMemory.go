package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeCurrentResCounterUsageForMemory(file *os.File, stats *types.Stats) (tab bool, err error) {
	// current res_counter usage for memory
	if e.rowsToPrint&KCurrentResCounterUsageForMemory == KCurrentResCounterUsageForMemory {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Usage)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KCurrentResCounterUsageForMemoryComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelCurrentResCounterUsageForMemory(file *os.File) (tab bool, err error) {
	// current res_counter usage for memory
	if e.rowsToPrint&KCurrentResCounterUsageForMemory == KCurrentResCounterUsageForMemory {
		_, err = file.Write([]byte("Current res_counter usage for memory"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KCurrentResCounterUsageForMemoryComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstCurrentResCounterUsageForMemory(file *os.File) (tab bool, err error) {
	// current res_counter usage for memory
	if e.rowsToPrint&KCurrentResCounterUsageForMemory == KCurrentResCounterUsageForMemory {
		_, err = file.Write([]byte("KCurrentResCounterUsageForMemory"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KCurrentResCounterUsageForMemoryComa != 0
	}

	return
}
