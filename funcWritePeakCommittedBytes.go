package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writePeakCommittedBytes(file *os.File, stats *types.Stats) (tab bool, err error) {
	// peak committed bytes
	if e.rowsToPrint&KPeakCommittedBytes == KPeakCommittedBytes {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.CommitPeak)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KPeakCommittedBytesComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelPeakCommittedBytes(file *os.File) (tab bool, err error) {
	// peak committed bytes
	if e.rowsToPrint&KPeakCommittedBytes == KPeakCommittedBytes {
		_, err = file.Write([]byte("Peak committed bytes"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KPeakCommittedBytesComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstPeakCommittedBytes(file *os.File) (tab bool, err error) {
	// peak committed bytes
	if e.rowsToPrint&KPeakCommittedBytes == KPeakCommittedBytes {
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
