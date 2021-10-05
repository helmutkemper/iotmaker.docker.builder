package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

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