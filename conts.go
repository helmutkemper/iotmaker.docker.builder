package iotmakerdockerbuilder

const (
	// KKiloByte
	//
	// English: 1024 Bytes multiplier
	//
	// Example:
	//   5 * KKiloByte = 5 KBytes
	//
	// Português: multiplicador de 1024 Bytes
	//
	// Exemplo:
	//   5 * KKiloByte = 5 KBytes
	KKiloByte = 1024

	// KMegaByte
	//
	// English: 1024 KBytes multiplier
	//
	// Example:
	//   5 * KMegaByte = 5 MBytes
	//
	// Português: multiplicador de 1024 KBytes
	//
	// Exemplo:
	//   5 * KMegaByte = 5 MBytes
	KMegaByte = 1024 * 1024

	// KGigaByte
	//
	// English: 1024 MBytes multiplier
	//
	// Example:
	//   5 * KGigaByte = 5 GBytes
	//
	// Português: multiplicador de 1024 MBytes
	//
	// Exemplo:
	//   5 * KGigaByte = 5 GBytes
	KGigaByte = 1024 * 1024 * 1024

	// KTeraByte (
	//
	// English: 1024 GBytes multiplier
	//
	// Example:
	//   5 * KTeraByte = 5 TBytes
	//
	// Português: multiplicador de 1024 GBytes
	//
	// Exemplo:
	//   5 * KTeraByte = 5 TBytes
	KTeraByte = 1024 * 1024 * 1024 * 1024

	// KAll
	//
	// English: Enable all values to log
	KAll = 0x7FFFFFFFFFFFFFF

	// KReadingTime
	//
	// English: Reading time
	KReadingTime     = 0b0000000000000000000000000000000000000000000000000000000000000001
	KReadingTimeComa = 0b0111111111111111111111111111111111111111111111111111111111111110

	KFilterLog     = 0b0000000000000000000000000000000000000000000000000000000000000010
	KFilterLogComa = 0b0111111111111111111111111111111111111111111111111111111111111100

	// KCurrentNumberOfOidsInTheCGroup
	//
	// English: Linux specific stats, not populated on Windows. Current is the number of pids in the cgroup
	KCurrentNumberOfOidsInTheCGroup     = 0b0000000000000000000000000000000000000000000000000000000000000100
	KCurrentNumberOfOidsInTheCGroupComa = 0b0111111111111111111111111111111111111111111111111111111111111000

	// KLimitOnTheNumberOfPidsInTheCGroup
	//
	// English: Linux specific stats, not populated on Windows. Limit is the hard limit on the number of pids in the cgroup. A "Limit" of 0 means that there is no limit.
	KLimitOnTheNumberOfPidsInTheCGroup     = 0b0000000000000000000000000000000000000000000000000000000000001000
	KLimitOnTheNumberOfPidsInTheCGroupComa = 0b0111111111111111111111111111111111111111111111111111111111110000

	// KTotalCPUTimeConsumed
	//
	// English: Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)
	KTotalCPUTimeConsumed     = 0b0000000000000000000000000000000000000000000000000000000000010000
	KTotalCPUTimeConsumedComa = 0b0111111111111111111111111111111111111111111111111111111111100000

	// KTotalCPUTimeConsumedPerCore
	//
	// English: Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)
	KTotalCPUTimeConsumedPerCore     = 0b0000000000000000000000000000000000000000000000000000000000100000
	KTotalCPUTimeConsumedPerCoreComa = 0b0111111111111111111111111111111111111111111111111111111111000000

	// KTimeSpentByTasksOfTheCGroupInKernelMode
	//
	// English: Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux). Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows.Not populated for Hyper-V Containers.)
	KTimeSpentByTasksOfTheCGroupInKernelMode     = 0b0000000000000000000000000000000000000000000000000000000001000000
	KTimeSpentByTasksOfTheCGroupInKernelModeComa = 0b0111111111111111111111111111111111111111111111111111111110000000

	// KTimeSpentByTasksOfTheCGroupInUserMode
	//
	// English: Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux). Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers)
	KTimeSpentByTasksOfTheCGroupInUserMode     = 0b0000000000000000000000000000000000000000000000000000000010000000
	KTimeSpentByTasksOfTheCGroupInUserModeComa = 0b0111111111111111111111111111111111111111111111111111111100000000

	// KSystemUsage
	//
	// English: System Usage. Linux only.
	KSystemUsage     = 0b0000000000000000000000000000000000000000000000000000000100000000
	KSystemUsageComa = 0b0111111111111111111111111111111111111111111111111111111000000000

	// KOnlineCPUs
	//
	// English: Online CPUs. Linux only.
	KOnlineCPUs     = 0b0000000000000000000000000000000000000000000000000000001000000000
	KOnlineCPUsComa = 0b0111111111111111111111111111111111111111111111111111110000000000

	// KNumberOfPeriodsWithThrottlingActive
	//
	// English: Throttling Data. Linux only. Number of periods with throttling active.
	KNumberOfPeriodsWithThrottlingActive     = 0b0000000000000000000000000000000000000000000000000000010000000000
	KNumberOfPeriodsWithThrottlingActiveComa = 0b0111111111111111111111111111111111111111111111111111100000000000

	// KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit
	//
	// English: Throttling Data. Linux only. Number of periods when the container hits its throttling limit.
	KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit     = 0b0000000000000000000000000000000000000000000000000000100000000000
	KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimitComa = 0b0111111111111111111111111111111111111111111111111111000000000000

	// KAggregateTimeTheContainerWasThrottledForInNanoseconds
	//
	// English: Throttling Data. Linux only. Aggregate time the container was throttled for in nanoseconds.
	KAggregateTimeTheContainerWasThrottledForInNanoseconds     = 0b0000000000000000000000000000000000000000000000000001000000000000
	KAggregateTimeTheContainerWasThrottledForInNanosecondsComa = 0b0111111111111111111111111111111111111111111111111110000000000000

	// KTotalPreCPUTimeConsumed
	//
	// English: Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows.
	KTotalPreCPUTimeConsumed     = 0b0000000000000000000000000000000000000000000000000010000000000000
	KTotalPreCPUTimeConsumedComa = 0b0111111111111111111111111111111111111111111111111100000000000000

	// KTotalPreCPUTimeConsumedPerCore
	//
	// English: Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows.
	KTotalPreCPUTimeConsumedPerCore     = 0b0000000000000000000000000000000000000000000000000100000000000000
	KTotalPreCPUTimeConsumedPerCoreComa = 0b0111111111111111111111111111111111111111111111111000000000000000

	// KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode
	//
	// English: Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux) - Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows - Not populated for Hyper-V Containers.)
	KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode     = 0b0000000000000000000000000000000000000000000000001000000000000000
	KTimeSpentByPreCPUTasksOfTheCGroupInKernelModeComa = 0b0111111111111111111111111111111111111111111111110000000000000000

	// KTimeSpentByPreCPUTasksOfTheCGroupInUserMode
	//
	// English: Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux) - Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers)
	KTimeSpentByPreCPUTasksOfTheCGroupInUserMode     = 0b0000000000000000000000000000000000000000000000010000000000000000
	KTimeSpentByPreCPUTasksOfTheCGroupInUserModeComa = 0b0111111111111111111111111111111111111111111111100000000000000000

	// KPreCPUSystemUsage
	//
	// English: System Usage. (Linux only)
	KPreCPUSystemUsage     = 0b0000000000000000000000000000000000000000000000100000000000000000
	KPreCPUSystemUsageComa = 0b0111111111111111111111111111111111111111111111000000000000000000

	// KOnlinePreCPUs
	//
	// English: Online CPUs. (Linux only)
	KOnlinePreCPUs     = 0b0000000000000000000000000000000000000000000001000000000000000000
	KOnlinePreCPUsComa = 0b0111111111111111111111111111111111111111111110000000000000000000

	// KAggregatePreCPUTimeTheContainerWasThrottled
	//
	// English: Throttling Data. (Linux only) - Aggregate time the container was throttled for in nanoseconds
	KAggregatePreCPUTimeTheContainerWasThrottled     = 0b0000000000000000000000000000000000000000000010000000000000000000
	KAggregatePreCPUTimeTheContainerWasThrottledComa = 0b0111111111111111111111111111111111111111111100000000000000000000

	// KNumberOfPeriodsWithPreCPUThrottlingActive
	//
	// English: Throttling Data. (Linux only) - Number of periods with throttling active
	KNumberOfPeriodsWithPreCPUThrottlingActive     = 0b0000000000000000000000000000000000000000000100000000000000000000
	KNumberOfPeriodsWithPreCPUThrottlingActiveComa = 0b0111111111111111111111111111111111111111111000000000000000000000

	// KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit
	//
	// English: Throttling Data. (Linux only) - Number of periods when the container hits its throttling limit.
	KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit     = 0b0000000000000000000000000000000000000000001000000000000000000000
	KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimitComa = 0b0111111111111111111111111111111111111111110000000000000000000000

	// KCurrentResCounterUsageForMemory
	//
	// English: Current res_counter usage for memory
	KCurrentResCounterUsageForMemory     = 0b0000000000000000000000000000000000000000010000000000000000000000
	KCurrentResCounterUsageForMemoryComa = 0b0111111111111111111111111111111111111111100000000000000000000000

	// KMaximumUsageEverRecorded
	//
	// English: Maximum usage ever recorded
	KMaximumUsageEverRecorded     = 0b0000000000000000000000000000000000000000100000000000000000000000
	KMaximumUsageEverRecordedComa = 0b0111111111111111111111111111111111111111000000000000000000000000

	// KNumberOfTimesMemoryUsageHitsLimits
	//
	// English: Number of times memory usage hits limits
	KNumberOfTimesMemoryUsageHitsLimits     = 0b0000000000000000000000000000000000000001000000000000000000000000
	KNumberOfTimesMemoryUsageHitsLimitsComa = 0b0111111111111111111111111111111111111110000000000000000000000000

	// KMemoryLimit
	//
	// English: Memory limit
	KMemoryLimit     = 0b0000000000000000000000000000000000000010000000000000000000000000
	KMemoryLimitComa = 0b0111111111111111111111111111111111111100000000000000000000000000

	// KCommittedBytes
	//
	// English: Committed bytes
	KCommittedBytes     = 0b0000000000000000000000000000000000000100000000000000000000000000
	KCommittedBytesComa = 0b0111111111111111111111111111111111111000000000000000000000000000

	// KPeakCommittedBytes
	//
	// English: Peak committed bytes
	KPeakCommittedBytes     = 0b0000000000000000000000000000000000001000000000000000000000000000
	KPeakCommittedBytesComa = 0b0111111111111111111111111111111111110000000000000000000000000000

	// KPrivateWorkingSet
	//
	// English: Private working set
	KPrivateWorkingSet     = 0b0000000000000000000000000000000000010000000000000000000000000000
	KPrivateWorkingSetComa = 0b0111111111111111111111111111111111100000000000000000000000000000

	KBlkioIoServiceBytesRecursive     = 0b0000000000000000000000000000000000100000000000000000000000000000
	KBlkioIoServiceBytesRecursiveComa = 0b0111111111111111111111111111111111000000000000000000000000000000

	KBlkioIoServicedRecursive     = 0b0000000000000000000000000000000001000000000000000000000000000000
	KBlkioIoServicedRecursiveComa = 0b0111111111111111111111111111111110000000000000000000000000000000

	KBlkioIoQueuedRecursive     = 0b0000000000000000000000000000000010000000000000000000000000000000
	KBlkioIoQueuedRecursiveComa = 0b0111111111111111111111111111111100000000000000000000000000000000

	KBlkioIoServiceTimeRecursive     = 0b0000000000000000000000000000000100000000000000000000000000000000
	KBlkioIoServiceTimeRecursiveComa = 0b0111111111111111111111111111111000000000000000000000000000000000

	KBlkioIoWaitTimeRecursive     = 0b0000000000000000000000000000001000000000000000000000000000000000
	KBlkioIoWaitTimeRecursiveComa = 0b0111111111111111111111111111110000000000000000000000000000000000

	KBlkioIoMergedRecursive     = 0b0000000000000000000000000000010000000000000000000000000000000000
	KBlkioIoMergedRecursiveComa = 0b0111111111111111111111111111100000000000000000000000000000000000

	KBlkioIoTimeRecursive     = 0b0000000000000000000000000000100000000000000000000000000000000000
	KBlkioIoTimeRecursiveComa = 0b0111111111111111111111111111000000000000000000000000000000000000

	KBlkioSectorsRecursive     = 0b0000000000000000000000000001000000000000000000000000000000000000
	KBlkioSectorsRecursiveComa = 0b0111111111111111111111111110000000000000000000000000000000000000

	// KMacOsLogWithAllCores
	//
	// English: Mac OS Log
	KMacOsLogWithAllCores = KReadingTime |
		KCurrentNumberOfOidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KSystemUsage |
		KOnlineCPUs |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KMemoryLimit |
		KBlkioIoServiceBytesRecursive | // não aparece no mac
		KBlkioIoServicedRecursive | // não aparece no mac
		KBlkioIoQueuedRecursive | // não aparece no mac
		KBlkioIoServiceTimeRecursive | // não aparece no mac
		KBlkioIoWaitTimeRecursive | // não aparece no mac
		KBlkioIoMergedRecursive | // não aparece no mac
		KBlkioIoTimeRecursive | // não aparece no mac
		KBlkioSectorsRecursive // não aparece no mac

	// KMacOsLog
	//
	// English: Mac OS Log
	KMacOsLog = KReadingTime |
		KCurrentNumberOfOidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KSystemUsage |
		KOnlineCPUs |
		KTotalPreCPUTimeConsumed |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KMemoryLimit |
		KBlkioIoServiceBytesRecursive | // não aparece no mac
		KBlkioIoServicedRecursive | // não aparece no mac
		KBlkioIoQueuedRecursive | // não aparece no mac
		KBlkioIoServiceTimeRecursive | // não aparece no mac
		KBlkioIoWaitTimeRecursive | // não aparece no mac
		KBlkioIoMergedRecursive | // não aparece no mac
		KBlkioIoTimeRecursive | // não aparece no mac
		KBlkioSectorsRecursive // não aparece no mac

	KWindows = KReadingTime |
		KCurrentNumberOfOidsInTheCGroup |
		KTotalCPUTimeConsumed |
		KTotalCPUTimeConsumedPerCore |
		KTimeSpentByTasksOfTheCGroupInKernelMode |
		KTimeSpentByTasksOfTheCGroupInUserMode |
		KSystemUsage |
		KOnlineCPUs |
		KTotalPreCPUTimeConsumed |
		KTotalPreCPUTimeConsumedPerCore |
		KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode |
		KTimeSpentByPreCPUTasksOfTheCGroupInUserMode |
		KPreCPUSystemUsage |
		KOnlinePreCPUs |
		KCurrentResCounterUsageForMemory |
		KMaximumUsageEverRecorded |
		KMemoryLimit
)
