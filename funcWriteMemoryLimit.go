package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeMemoryLimit(file *os.File, stats *types.Stats) (tab bool, err error) {
	if e.rowsToPrint&KMemoryLimit == KMemoryLimit {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Limit)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KMemoryLimitComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelMemoryLimit(file *os.File) (tab bool, err error) {
	if e.rowsToPrint&KMemoryLimit == KMemoryLimit {
		_, err = file.Write([]byte("Memory limit"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KMemoryLimitComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstMemoryLimit(file *os.File) (tab bool, err error) {
	if e.rowsToPrint&KMemoryLimit == KMemoryLimit {
		_, err = file.Write([]byte("KMemoryLimit"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KMemoryLimitComa != 0
	}

	return
}
