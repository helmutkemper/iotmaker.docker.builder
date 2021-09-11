package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeTimeSpentByPreCPUTasksOfTheCGroupInUserMode(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in user mode (Linux).
	// Time spent by all container processes in user mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers
	if makeLabel == true && e.logFlags&KTimeSpentByPreCPUTasksOfTheCGroupInUserMode == KTimeSpentByPreCPUTasksOfTheCGroupInUserMode {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux) - Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers)\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTimeSpentByPreCPUTasksOfTheCGroupInUserMode == KTimeSpentByPreCPUTasksOfTheCGroupInUserMode {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.CPUUsage.UsageInUsermode)))
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
