package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeSystemUsage(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// System Usage. Linux only.
	if makeLabel == true && e.logFlags&KSystemUsage == KSystemUsage {
		_, err = file.Write([]byte("System Usage. Linux only.\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KSystemUsage == KSystemUsage {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.SystemUsage)))
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
