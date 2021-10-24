package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeConstTotalCPUTimeConsumed(file *os.File) (tab bool, err error) {
	// Total CPU time consumed.
	// Units: nanoseconds (Linux)
	// Units: 100's of nanoseconds (Windows)
	if e.rowsToPrint&KLogColumnTotalCPUTimeConsumed == KLogColumnTotalCPUTimeConsumed {
		_, err = file.Write([]byte("KTotalCPUTimeConsumed"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KTotalCPUTimeConsumedComa != 0
	}

	return
}
