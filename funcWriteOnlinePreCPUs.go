package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeOnlinePreCPUs(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Online CPUs. Linux only.
	if makeLabel == true && e.logFlags&KOnlinePreCPUs == KOnlinePreCPUs {
		_, err = file.Write([]byte("Online CPUs. (Linux only)\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KOnlinePreCPUs == KOnlinePreCPUs {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.OnlineCPUs)))
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
