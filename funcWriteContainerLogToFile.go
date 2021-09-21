package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"io/fs"
	"log"
	"os"
)

// writeContainerLogToFile
//
// Português: Escreve um arquivo csv com dados capturados da saída padrão do container e dados estatísticos do container
func (e *ContainerBuilder) writeContainerLogToFile(path string, lineList [][]byte) (err error) {
	if path == "" {
		return
	}

	if lineList == nil {
		return
	}

	var makeLabel = false
	_, err = os.Stat(path)
	if err != nil {
		makeLabel = true
	}

	var file *os.File
	file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
		}
	}(file)

	var stats = types.Stats{}
	stats, err = e.ContainerStatisticsOneShot()
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if makeLabel == true && e.csvConstHeader == true {
		err = e.writeContainerConstToFile(file, &stats)
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte(e.csvRowSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	if makeLabel == true {
		err = e.writeContainerLabelToFile(file, &stats)
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte(e.csvRowSeparator))
		if err != nil {
			log.Printf("writeContainerLogToFile().error: %v", err.Error())
			util.TraceToLog()
			return
		}
	}

	err = e.writeContainerStatsToFile(file, &stats, &lineList)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	_, err = file.Write([]byte(e.csvRowSeparator))
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	return
}

func (e *ContainerBuilder) writeContainerStatsToFile(file *os.File, stats *types.Stats, lineList *[][]byte) (err error) {
	var tab bool

	// time ok
	tab, err = e.writeReadingTime(file, stats)
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

	tab, err = e.writeFilterIntoLog(file, e.chaos.filterLog, lineList)
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

	tab, err = e.writeCurrentNumberOfOidsInTheCGroup(file, stats)
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

	tab, err = e.writeLimitOnTheNumberOfPidsInTheCGroup(file, stats)
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

	tab, err = e.writeTotalCPUTimeConsumed(file, stats)
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

	tab, err = e.writeTotalCPUTimeConsumedPerCore(file, stats)
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

	tab, err = e.writeTimeSpentByTasksOfTheCGroupInKernelMode(file, stats)
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

	tab, err = e.writeTimeSpentByTasksOfTheCGroupInUserMode(file, stats)
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

	tab, err = e.writeSystemUsage(file, stats)
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

	tab, err = e.writeOnlineCPUs(file, stats)
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

	tab, err = e.writeNumberOfPeriodsWithThrottlingActive(file, stats)
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

	tab, err = e.writeNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit(file, stats)
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

	tab, err = e.writeAggregateTimeTheContainerWasThrottledForInNanoseconds(file, stats)
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

	tab, err = e.writeTotalPreCPUTimeConsumed(file, stats)
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

	tab, err = e.writeTotalPreCPUTimeConsumedPerCore(file, stats)
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

	tab, err = e.writeTimeSpentByPreCPUTasksOfTheCGroupInKernelMode(file, stats)
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

	tab, err = e.writeTimeSpentByPreCPUTasksOfTheCGroupInUserMode(file, stats)
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

	tab, err = e.writePreCPUSystemUsage(file, stats)
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

	tab, err = e.writeOnlinePreCPUs(file, stats)
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

	tab, err = e.writeAggregatePreCPUTimeTheContainerWasThrottled(file, stats)
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

	tab, err = e.writeNumberOfPeriodsWithPreCPUThrottlingActive(file, stats)
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

	tab, err = e.writeNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit(file, stats)
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

	tab, err = e.writeCurrentResCounterUsageForMemory(file, stats)
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

	tab, err = e.writeMaximumUsageEverRecorded(file, stats)
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

	tab, err = e.writeNumberOfTimesMemoryUsageHitsLimits(file, stats)
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

	tab, err = e.writeMemoryLimit(file, stats)
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

	tab, err = e.writeCommittedBytes(file, stats)
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

	tab, err = e.writePeakCommittedBytes(file, stats)
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

	tab, err = e.writePrivateWorkingSet(file, stats)
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

	tab, err = e.writeBlkioIoServiceBytesRecursive(file, stats)
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

	tab, err = e.writeBlkioIoServicedRecursive(file, stats)
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

	tab, err = e.writeBlkioIoQueuedRecursive(file, stats)
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

	tab, err = e.writeBlkioIoServiceTimeRecursive(file, stats)
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

	tab, err = e.writeBlkioIoWaitTimeRecursive(file, stats)
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

	tab, err = e.writeBlkioIoMergedRecursive(file, stats)
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

	tab, err = e.writeBlkioIoTimeRecursive(file, stats)
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

	_, err = e.writeBlkioSectorsRecursive(file, stats)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	return
}
