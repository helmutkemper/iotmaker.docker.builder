package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeTimeSpentByTasksOfTheCGroupInKernelMode(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	if makeLabel == true && e.logFlags&KTimeSpentByTasksOfTheCGroupInKernelMode == KTimeSpentByTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux). Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows.Not populated for Hyper-V Containers.).\t"))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTimeSpentByTasksOfTheCGroupInKernelMode == KTimeSpentByTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.CPUUsage.UsageInKernelmode)))
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
