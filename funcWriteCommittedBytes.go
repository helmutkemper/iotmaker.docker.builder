package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeCommittedBytes(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// committed bytes
	if makeLabel == true && e.logFlags&KCommittedBytes == KCommittedBytes {
		_, err = file.Write([]byte("Committed bytes\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KCommittedBytes == KCommittedBytes {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Commit)))
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
