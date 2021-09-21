package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeLabelTimeSpentByPreCPUTasksOfTheCGroupInKernelMode(file *os.File) (tab bool, err error) {
	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	if e.rowsToPrint&KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode == KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux) - Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows - Not populated for Hyper-V Containers.)"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KTimeSpentByPreCPUTasksOfTheCGroupInKernelModeComa != 0
	}

	return
}
