package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeCurrentResCounterUsageForMemory(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// current res_counter usage for memory
	if makeLabel == true && e.logFlags&KCurrentResCounterUsageForMemory == KCurrentResCounterUsageForMemory {
		_, err = file.Write([]byte("Current res_counter usage for memory\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KCurrentResCounterUsageForMemory == KCurrentResCounterUsageForMemory {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Usage)))
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
