package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeurrentNumberOfOidsInTheCGroup(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Linux specific stats, not populated on Windows.
	// Current is the number of pids in the cgroup
	if makeLabel == true && e.logFlags&KCurrentNumberOfOidsInTheCGroup == KCurrentNumberOfOidsInTheCGroup {
		_, err = file.Write([]byte("Linux specific stats, not populated on Windows. Current is the number of pids in the cgroup\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KCurrentNumberOfOidsInTheCGroup == KCurrentNumberOfOidsInTheCGroup {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PidsStats.Current)))
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
