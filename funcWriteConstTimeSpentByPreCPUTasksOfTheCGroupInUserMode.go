package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeConstTimeSpentByPreCPUTasksOfTheCGroupInUserMode(file *os.File) (tab bool, err error) {
	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in user mode (Linux).
	// Time spent by all container processes in user mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers
	if e.rowsToPrint&KTimeSpentByPreCPUTasksOfTheCGroupInUserMode == KTimeSpentByPreCPUTasksOfTheCGroupInUserMode {
		_, err = file.Write([]byte("KTimeSpentByPreCPUTasksOfTheCGroupInUserMode"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KTimeSpentByPreCPUTasksOfTheCGroupInUserModeComa != 0
	}

	return
}
