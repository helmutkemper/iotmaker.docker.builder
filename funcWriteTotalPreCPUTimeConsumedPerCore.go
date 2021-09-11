package iotmakerdockerbuilder

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeTotalPreCPUTimeConsumedPerCore(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	if makeLabel == true && e.logFlags&KTotalPreCPUTimeConsumedPerCore == KTotalPreCPUTimeConsumedPerCore {
		for cpuNumber := 0; cpuNumber != e.logCpus; cpuNumber += 1 {
			_, err = file.Write([]byte(fmt.Sprintf("Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows. CPU: %v\t", cpuNumber)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KTotalPreCPUTimeConsumedPerCore == KTotalPreCPUTimeConsumedPerCore {
		// CPU Usage. Linux and Windows.
		// Total CPU time consumed per core (Linux). Not used on Windows.
		// Units: nanoseconds.
		if e.logCpus != 0 && len(stats.CPUStats.CPUUsage.PercpuUsage) == 0 {
			for i := 0; i != e.logCpus; i += 1 {
				_, err = file.Write([]byte{0x30})
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
		}

		for _, cpuTime := range stats.PreCPUStats.CPUUsage.PercpuUsage {
			_, err = file.Write([]byte(fmt.Sprintf("%v", cpuTime)))
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
	}

	return
}
