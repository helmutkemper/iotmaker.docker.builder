package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeCurrentNumberOfOidsInTheCGroup(file *os.File, stats *types.Stats) (tab bool, err error) {
	// Linux specific stats, not populated on Windows.
	// Current is the number of pids in the cgroup
	if e.rowsToPrint&KCurrentNumberOfOidsInTheCGroup == KCurrentNumberOfOidsInTheCGroup {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PidsStats.Current)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KCurrentNumberOfOidsInTheCGroupComa != 0
	}

	return
}
