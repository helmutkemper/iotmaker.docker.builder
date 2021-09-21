package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writePrivateWorkingSet(file *os.File, stats *types.Stats) (tab bool, err error) {
	// private working set
	if e.rowsToPrint&KPrivateWorkingSet == KPrivateWorkingSet {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.PrivateWorkingSet)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KPrivateWorkingSetComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelPrivateWorkingSet(file *os.File) (tab bool, err error) {
	// private working set
	if e.rowsToPrint&KPrivateWorkingSet == KPrivateWorkingSet {
		_, err = file.Write([]byte("Private working set"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KPrivateWorkingSetComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstPrivateWorkingSet(file *os.File) (tab bool, err error) {
	// private working set
	if e.rowsToPrint&KPrivateWorkingSet == KPrivateWorkingSet {
		_, err = file.Write([]byte("KPrivateWorkingSet"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KPrivateWorkingSetComa != 0
	}

	return
}
