package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeConstSystemUsage(file *os.File) (tab bool, err error) {
	// System Usage. Linux only.
	if e.rowsToPrint&KSystemUsage == KSystemUsage {
		_, err = file.Write([]byte("KSystemUsage"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KSystemUsageComa != 0
	}

	return
}
