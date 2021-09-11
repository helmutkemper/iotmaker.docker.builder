package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writePrivateWorkingSet(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// private working set
	if makeLabel == true && e.logFlags&KPrivateWorkingSet == KPrivateWorkingSet {
		_, err = file.Write([]byte("Private working set\n"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KPrivateWorkingSet == KPrivateWorkingSet {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.PrivateWorkingSet)))
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
