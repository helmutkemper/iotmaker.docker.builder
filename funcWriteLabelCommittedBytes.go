package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeLabelCommittedBytes(file *os.File) (tab bool, err error) {
	// committed bytes
	if e.rowsToPrint&KCommittedBytes == KCommittedBytes {
		_, err = file.Write([]byte("Committed bytes"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KCommittedBytesComa != 0
	}

	return
}
