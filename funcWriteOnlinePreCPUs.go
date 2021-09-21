package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeOnlinePreCPUs(file *os.File, stats *types.Stats) (tab bool, err error) {
	// Online CPUs. Linux only.
	if e.rowsToPrint&KOnlinePreCPUs == KOnlinePreCPUs {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.OnlineCPUs)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KOnlinePreCPUsComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelOnlinePreCPUs(file *os.File) (tab bool, err error) {
	// Online CPUs. Linux only.
	if e.rowsToPrint&KOnlinePreCPUs == KOnlinePreCPUs {
		_, err = file.Write([]byte("Online CPUs. (Linux only)"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KOnlinePreCPUsComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstOnlinePreCPUs(file *os.File) (tab bool, err error) {
	// Online CPUs. Linux only.
	if e.rowsToPrint&KOnlinePreCPUs == KOnlinePreCPUs {
		_, err = file.Write([]byte("KOnlinePreCPUs"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KOnlinePreCPUsComa != 0
	}

	return
}
