package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeConstPeakCommittedBytes(file *os.File) (tab bool, err error) {
	// peak committed bytes
	if e.rowsToPrint&KLogColumnPeakCommittedBytes == KLogColumnPeakCommittedBytes {
		_, err = file.Write([]byte("KPeakCommittedBytes"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KPeakCommittedBytesComa != 0
	}

	return
}
