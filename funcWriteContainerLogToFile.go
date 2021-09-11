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

	// time
	err = e.writeTime(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeFilterIntoLog(file, e.chaos.filterLog, &lineList, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioIoServiceBytesRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioIoServicedRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioIoQueuedRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioIoServiceTimeRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioIoWaitTimeRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioIoMergedRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioIoTimeRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeBlkioSectorsRecursive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeurrentNumberOfOidsInTheCGroup(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeLimitOnTheNumberOfPidsInTheCGroup(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeTotalCPUTimeConsumed(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	if len(stats.CPUStats.CPUUsage.PercpuUsage) != 0 {
		e.logCpus = len(stats.CPUStats.CPUUsage.PercpuUsage)
	}

	err = e.writeTotalCPUTimeConsumedPerCore(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeTimeSpentByTasksOfTheCGroupInKernelMode(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeTimeSpentByTasksOfTheCGroupInUserMode(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeSystemUsage(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeOnlineCPUs(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeNumberOfPeriodsWithThrottlingActive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeAggregateTimeTheContainerWasThrottledForInNanoseconds(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeTotalPreCPUTimeConsumed(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeTotalPreCPUTimeConsumedPerCore(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeTimeSpentByPreCPUTasksOfTheCGroupInKernelMode(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeTimeSpentByPreCPUTasksOfTheCGroupInUserMode(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writePreCPUSystemUsage(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeOnlinePreCPUs(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeAggregatePreCPUTimeTheContainerWasThrottled(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeNumberOfPeriodsWithPreCPUThrottlingActive(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeCurrentResCounterUsageForMemory(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeMaximumUsageEverRecorded(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeNumberOfTimesMemoryUsageHitsLimits(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeMemoryLimit(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writeCommittedBytes(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writePeakCommittedBytes(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	err = e.writePrivateWorkingSet(file, &stats, makeLabel)
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	_, err = file.Write([]byte("\n"))
	if err != nil {
		log.Printf("writeContainerLogToFile().error: %v", err.Error())
		util.TraceToLog()
		return
	}

	return
}
