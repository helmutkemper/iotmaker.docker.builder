package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
)

func (e *ContainerBuilder) writeContainerLabelToFile(file *os.File, stats *types.Stats) (err error) {
	var tab bool

	// time ok
	tab, err = e.writeLabelReadingTime(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	for _, v := range e.chaos.filterLog {
		if v.Label != "" {
			tab = true
			break
		}
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelFilterIntoLog(file, e.chaos.filterLog)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelCurrentNumberOfOidsInTheCGroup(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelLimitOnTheNumberOfPidsInTheCGroup(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelTotalCPUTimeConsumed(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	if len(stats.CPUStats.CPUUsage.PercpuUsage) != 0 {
		e.logCpus = len(stats.CPUStats.CPUUsage.PercpuUsage)
	}

	tab, err = e.writeLabelTotalCPUTimeConsumedPerCore(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelTimeSpentByTasksOfTheCGroupInKernelMode(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelTimeSpentByTasksOfTheCGroupInUserMode(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelSystemUsage(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelOnlineCPUs(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelNumberOfPeriodsWithThrottlingActive(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelAggregateTimeTheContainerWasThrottledForInNanoseconds(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelTotalPreCPUTimeConsumed(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelTotalPreCPUTimeConsumedPerCore(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelTimeSpentByPreCPUTasksOfTheCGroupInKernelMode(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelTimeSpentByPreCPUTasksOfTheCGroupInUserMode(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelPreCPUSystemUsage(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelOnlinePreCPUs(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelAggregatePreCPUTimeTheContainerWasThrottled(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelNumberOfPeriodsWithPreCPUThrottlingActive(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelCurrentResCounterUsageForMemory(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelMaximumUsageEverRecorded(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelNumberOfTimesMemoryUsageHitsLimits(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelMemoryLimit(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelCommittedBytes(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelPeakCommittedBytes(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelPrivateWorkingSet(file)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelBlkioIoServiceBytesRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelBlkioIoServicedRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelBlkioIoQueuedRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelBlkioIoServiceTimeRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelBlkioIoWaitTimeRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelBlkioIoMergedRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	tab, err = e.writeLabelBlkioIoTimeRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if tab == true {
		_, err = file.Write([]byte(e.csvValueSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	_, err = e.writeLabelBlkioSectorsRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	return
}
