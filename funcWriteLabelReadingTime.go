package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeReadingTime(file *os.File, stats *types.Stats) (tab bool, err error) {
	if e.rowsToPrint&KReadingTime == KReadingTime {
		_, err = file.Write([]byte(stats.Read.String()))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
		tab = e.rowsToPrint&KReadingTimeComa != 0
	}

	return
}

func (e *ContainerBuilder) writeLabelReadingTime(file *os.File) (tab bool, err error) {
	if e.rowsToPrint&KReadingTime == KReadingTime {
		_, err = file.Write([]byte("Reading time"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
		tab = e.rowsToPrint&KReadingTimeComa != 0
	}

	return
}

func (e *ContainerBuilder) writeConstReadingTime(file *os.File) (tab bool, err error) {
	if e.rowsToPrint&KReadingTime == KReadingTime {
		_, err = file.Write([]byte("KReadingTime"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KReadingTimeComa != 0
	}

	return
}
