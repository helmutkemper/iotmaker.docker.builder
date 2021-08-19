package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
)

func (e *ContainerBuilder) ContainerStatisticsOneShot() (
	statsRet Stats,
	err error,
) {
	var statsRetTypes types.Stats
	statsRetTypes, err = e.dockerSys.ContainerStatisticsOneShot(e.containerID)
	if err != nil {
		return
	}

	statsRet.Read = statsRetTypes.Read
	statsRet.PreRead = statsRetTypes.PreRead
	statsRet.PidsStats = PidsStats(statsRetTypes.PidsStats)

	statsRet.BlkioStats.IoServiceBytesRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.IoServiceBytesRecursive))
	for k, v := range statsRetTypes.BlkioStats.IoServiceBytesRecursive {
		statsRet.BlkioStats.IoServiceBytesRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.BlkioStats.IoServicedRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.IoServicedRecursive))
	for k, v := range statsRetTypes.BlkioStats.IoServicedRecursive {
		statsRet.BlkioStats.IoServicedRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.BlkioStats.IoQueuedRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.IoQueuedRecursive))
	for k, v := range statsRetTypes.BlkioStats.IoQueuedRecursive {
		statsRet.BlkioStats.IoQueuedRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.BlkioStats.IoServiceTimeRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.IoServiceTimeRecursive))
	for k, v := range statsRetTypes.BlkioStats.IoServiceTimeRecursive {
		statsRet.BlkioStats.IoServiceTimeRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.BlkioStats.IoWaitTimeRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.IoWaitTimeRecursive))
	for k, v := range statsRetTypes.BlkioStats.IoWaitTimeRecursive {
		statsRet.BlkioStats.IoWaitTimeRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.BlkioStats.IoMergedRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.IoMergedRecursive))
	for k, v := range statsRetTypes.BlkioStats.IoMergedRecursive {
		statsRet.BlkioStats.IoMergedRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.BlkioStats.IoTimeRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.IoTimeRecursive))
	for k, v := range statsRetTypes.BlkioStats.IoTimeRecursive {
		statsRet.BlkioStats.IoTimeRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.BlkioStats.SectorsRecursive = make([]BlkioStatEntry, len(statsRetTypes.BlkioStats.SectorsRecursive))
	for k, v := range statsRetTypes.BlkioStats.SectorsRecursive {
		statsRet.BlkioStats.SectorsRecursive[k] = BlkioStatEntry(v)
	}

	statsRet.NumProcs = statsRetTypes.NumProcs
	statsRet.StorageStats = StorageStats(statsRetTypes.StorageStats)

	statsRet.CPUStats.CPUUsage = CPUUsage(statsRetTypes.CPUStats.CPUUsage)
	statsRet.CPUStats.SystemUsage = statsRetTypes.CPUStats.SystemUsage
	statsRet.CPUStats.OnlineCPUs = statsRetTypes.CPUStats.OnlineCPUs
	statsRet.CPUStats.ThrottlingData = ThrottlingData(statsRetTypes.CPUStats.ThrottlingData)

	statsRet.PreCPUStats.CPUUsage = CPUUsage(statsRetTypes.PreCPUStats.CPUUsage)
	statsRet.PreCPUStats.SystemUsage = statsRetTypes.PreCPUStats.SystemUsage
	statsRet.PreCPUStats.OnlineCPUs = statsRetTypes.PreCPUStats.OnlineCPUs
	statsRet.PreCPUStats.ThrottlingData = ThrottlingData(statsRetTypes.PreCPUStats.ThrottlingData)

	statsRet.MemoryStats.Usage = statsRetTypes.MemoryStats.Usage
	statsRet.MemoryStats.MaxUsage = statsRetTypes.MemoryStats.MaxUsage
	statsRet.MemoryStats.Stats = statsRetTypes.MemoryStats.Stats
	statsRet.MemoryStats.Failcnt = statsRetTypes.MemoryStats.Failcnt
	statsRet.MemoryStats.Limit = statsRetTypes.MemoryStats.Limit
	statsRet.MemoryStats.Commit = statsRetTypes.MemoryStats.Commit
	statsRet.MemoryStats.CommitPeak = statsRetTypes.MemoryStats.CommitPeak
	statsRet.MemoryStats.PrivateWorkingSet = statsRetTypes.MemoryStats.PrivateWorkingSet

	return
}
