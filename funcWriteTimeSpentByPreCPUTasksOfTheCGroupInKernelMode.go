package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeTimeSpentByPreCPUTasksOfTheCGroupInKernelMode(file *os.File, stats *types.Stats) (tab bool, err error) {
	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	if e.rowsToPrint&KLogColumnTimeSpentByPreCPUTasksOfTheCGroupInKernelMode == KLogColumnTimeSpentByPreCPUTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.CPUUsage.UsageInKernelmode)))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		tab = e.rowsToPrint&KTimeSpentByPreCPUTasksOfTheCGroupInKernelModeComa != 0
	}

	return
}
