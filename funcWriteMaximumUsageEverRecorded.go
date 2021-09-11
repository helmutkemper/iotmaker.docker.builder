package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeMaximumUsageEverRecorded(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// maximum usage ever recorded.
	if makeLabel == true && e.logFlags&KMaximumUsageEverRecorded == KMaximumUsageEverRecorded {
		_, err = file.Write([]byte("Maximum usage ever recorded.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KMaximumUsageEverRecorded == KMaximumUsageEverRecorded {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.MaxUsage)))
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
