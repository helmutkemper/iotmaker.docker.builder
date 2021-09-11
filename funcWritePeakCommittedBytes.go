package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writePeakCommittedBytes(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// peak committed bytes
	if makeLabel == true && e.logFlags&KPeakCommittedBytes == KPeakCommittedBytes {
		_, err = file.Write([]byte("Peak committed bytes\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KPeakCommittedBytes == KPeakCommittedBytes {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.CommitPeak)))
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
