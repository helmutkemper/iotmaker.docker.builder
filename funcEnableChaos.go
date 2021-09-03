package iotmakerdockerbuilder

import (
	"bytes"
	"fmt"
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
	"io/fs"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

const (

	// KAll
	//
	// English: Enable all values to log
	KAll = 0x8FFFFFFFFFFFFFF

	// KReadingTime
	//
	// English: Reading time
	KReadingTime = 0b0000000000000000000000000000000000000000000000000000000000000001

	// KCurrentNumberOfOidsInTheCGroup
	//
	// English: Linux specific stats, not populated on Windows. Current is the number of pids in the cgroup
	KCurrentNumberOfOidsInTheCGroup = 0b0000000000000000000000000000000000000000000000000000000000000010

	// KLimitOnTheNumberOfPidsInTheCGroup
	//
	// English: Linux specific stats, not populated on Windows. Limit is the hard limit on the number of pids in the cgroup. A "Limit" of 0 means that there is no limit.
	KLimitOnTheNumberOfPidsInTheCGroup = 0b0000000000000000000000000000000000000000000000000000000000000100

	// KTotalCPUTimeConsumed
	//
	// English: Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)
	KTotalCPUTimeConsumed = 0b0000000000000000000000000000000000000000000000000000000000001000

	// KTotalCPUTimeConsumedPerCore
	//
	// English: Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)
	KTotalCPUTimeConsumedPerCore = 0b0000000000000000000000000000000000000000000000000000000000010000

	// KTimeSpentByTasksOfTheCGroupInKernelMode
	//
	// English: Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux). Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows.Not populated for Hyper-V Containers.)
	KTimeSpentByTasksOfTheCGroupInKernelMode = 0b0000000000000000000000000000000000000000000000000000000000100000

	// KTimeSpentByTasksOfTheCGroupInUserMode
	//
	// English: Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux). Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers)
	KTimeSpentByTasksOfTheCGroupInUserMode = 0b0000000000000000000000000000000000000000000000000000000001000000

	// KSystemUsage
	//
	// English: System Usage. Linux only.
	KSystemUsage = 0b0000000000000000000000000000000000000000000000000000000010000000

	// KOnlineCPUs
	//
	// English: Online CPUs. Linux only.
	KOnlineCPUs = 0b0000000000000000000000000000000000000000000000000000000100000000

	// KNumberOfPeriodsWithThrottlingActive
	//
	// English: Throttling Data. Linux only. Number of periods with throttling active.
	KNumberOfPeriodsWithThrottlingActive = 0b0000000000000000000000000000000000000000000000000000001000000000

	// KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit
	//
	// English: Throttling Data. Linux only. Number of periods when the container hits its throttling limit.
	KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit = 0b0000000000000000000000000000000000000000000000000000010000000000

	// KAggregateTimeTheContainerWasThrottledForInNanoseconds
	//
	// English: Throttling Data. Linux only. Aggregate time the container was throttled for in nanoseconds.
	KAggregateTimeTheContainerWasThrottledForInNanoseconds = 0b0000000000000000000000000000000000000000000000000000100000000000

	// KTotalPreCPUTimeConsumed
	//
	// English: Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows.
	KTotalPreCPUTimeConsumed = 0b0000000000000000000000000000000000000000000000000001000000000000

	// KTotalPreCPUTimeConsumedPerCore
	//
	// English: Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows.
	KTotalPreCPUTimeConsumedPerCore = 0b0000000000000000000000000000000000000000000000000010000000000000

	// KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode
	//
	// English: Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux) - Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows - Not populated for Hyper-V Containers.)
	KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode = 0b0000000000000000000000000000000000000000000000000100000000000000

	// KTimeSpentByPreCPUTasksOfTheCGroupInUserMode
	//
	// English: Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux) - Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers)
	KTimeSpentByPreCPUTasksOfTheCGroupInUserMode = 0b0000000000000000000000000000000000000000000000001000000000000000

	// KPreCPUSystemUsage
	//
	// English: System Usage. (Linux only)
	KPreCPUSystemUsage = 0b0000000000000000000000000000000000000000000000010000000000000000

	// KOnlinePreCPUs
	//
	// English: Online CPUs. (Linux only)
	KOnlinePreCPUs = 0b0000000000000000000000000000000000000000000000100000000000000000

	// KAggregatePreCPUTimeTheContainerWasThrottled
	//
	// English: Throttling Data. (Linux only) - Aggregate time the container was throttled for in nanoseconds
	KAggregatePreCPUTimeTheContainerWasThrottled = 0b0000000000000000000000000000000000000000000001000000000000000000

	// KNumberOfPeriodsWithPreCPUThrottlingActive
	//
	// English: Throttling Data. (Linux only) - Number of periods with throttling active
	KNumberOfPeriodsWithPreCPUThrottlingActive = 0b0000000000000000000000000000000000000000000010000000000000000000

	// KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit
	//
	// English: Throttling Data. (Linux only) - Number of periods when the container hits its throttling limit.
	KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit = 0b0000000000000000000000000000000000000000000100000000000000000000

	// KCurrentResCounterUsageForMemory
	//
	// English: Current res_counter usage for memory
	KCurrentResCounterUsageForMemory = 0b0000000000000000000000000000000000000000001000000000000000000000

	// KMaximumUsageEverRecorded
	//
	// English: Maximum usage ever recorded
	KMaximumUsageEverRecorded = 0b0000000000000000000000000000000000000000010000000000000000000000

	// KNumberOfTimesMemoryUsageHitsLimits
	//
	// English: Number of times memory usage hits limits
	KNumberOfTimesMemoryUsageHitsLimits = 0b0000000000000000000000000000000000000000100000000000000000000000

	// KMemoryLimit
	//
	// English: Memory limit
	KMemoryLimit = 0b0000000000000000000000000000000000000001000000000000000000000000

	// KCommittedBytes
	//
	// English: Committed bytes
	KCommittedBytes = 0b0000000000000000000000000000000000000010000000000000000000000000

	// KPeakCommittedBytes
	//
	// English: Peak committed bytes
	KPeakCommittedBytes = 0b0000000000000000000000000000000000000100000000000000000000000000

	// KPrivateWorkingSet
	//
	// English: Private working set
	KPrivateWorkingSet = 0b0000000000000000000000000000000000001000000000000000000000000000

	KBlkioIoServiceBytesRecursive = 0b0000000000000000000000000000000000010000000000000000000000000000
	KBlkioIoServicedRecursive     = 0b0000000000000000000000000000000000100000000000000000000000000000
	KBlkioIoQueuedRecursive       = 0b0000000000000000000000000000000001000000000000000000000000000000
	KBlkioIoServiceTimeRecursive  = 0b0000000000000000000000000000000010000000000000000000000000000000
	KBlkioIoWaitTimeRecursive     = 0b0000000000000000000000000000000100000000000000000000000000000000
	KBlkioIoMergedRecursive       = 0b0000000000000000000000000000001000000000000000000000000000000000
	KBlkioIoTimeRecursive         = 0b0000000000000000000000000000010000000000000000000000000000000000
	KBlkioSectorsRecursive        = 0b0000000000000000000000000000100000000000000000000000000000000000

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
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive |
		KBlkioIoServiceTimeRecursive |
		KBlkioIoWaitTimeRecursive |
		KBlkioIoMergedRecursive |
		KBlkioIoTimeRecursive |
		KBlkioSectorsRecursive

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
		KBlkioIoServiceBytesRecursive |
		KBlkioIoServicedRecursive |
		KBlkioIoQueuedRecursive |
		KBlkioIoServiceTimeRecursive |
		KBlkioIoWaitTimeRecursive |
		KBlkioIoMergedRecursive |
		KBlkioIoTimeRecursive |
		KBlkioSectorsRecursive
)

func (e *ContainerBuilder) EnableChaos(enable bool) {
	e.chaos.enableChaos = enable
}

func (e ContainerBuilder) traceCodeLine() (file string, line int) {
	_, file, line, _ = runtime.Caller(1)
	return
}

func (e ContainerBuilder) logsCleaner(logs []byte) [][]byte {
	logs = bytes.ReplaceAll(logs, []byte("\r"), []byte(""))
	return bytes.Split(logs, []byte("\n"))
}

func (e *ContainerBuilder) selectBetweenMaxAndMin(max, min time.Duration) (selected time.Duration) {
	randValue := e.getRandSeed().Int63n(int64(max)-int64(min)) + int64(min)
	return time.Duration(randValue)
}

func (e *ContainerBuilder) logsSearchAndReplaceIntoText(lineList [][]byte, configuration []LogFilter) (line []byte, found bool) {
	var err error
	if configuration == nil {
		return
	}

	for logLine := len(lineList) - 1; logLine >= 0; logLine -= 1 {

		for filterLine := 0; filterLine != len(configuration); filterLine += 1 {
			line = lineList[logLine]
			if bytes.Contains(line, []byte(configuration[filterLine].Match)) == true {

				if configuration[filterLine].Filter != "" {

					var re *regexp.Regexp
					re, err = regexp.Compile(configuration[filterLine].Filter)
					if err != nil {
						util.TraceToLog()
						log.Printf("regexp.Compile().error: %v", err)
						log.Printf("regexp.Compile().error.filter: %v", configuration[filterLine].Filter)
						continue
					}

					line = re.ReplaceAll(lineList[logLine], []byte("${valueToGet}"))

					if configuration[filterLine].Search != "" {
						re, err = regexp.Compile(configuration[filterLine].Search)
						if err != nil {
							util.TraceToLog()
							log.Printf("regexp.Compile().error: %v", err)
							log.Printf("regexp.Compile().error.filter: %v", configuration[filterLine].Search)
							continue
						}

						line = re.ReplaceAll(line, []byte(configuration[filterLine].Replace))
					}
				}

				found = true
				return
			}
		}
	}

	return
}

func (e *ContainerBuilder) managerChaos() {
	var err error
	var logs []byte
	var lineList [][]byte
	var line []byte
	var found bool
	var timeToNextEvent time.Duration
	var probality float64
	var lineNumber int
	var event Event

	var inspect iotmakerdocker.ContainerInspect

	probality = e.getProbalityNumber()

	inspect, err = e.ContainerInspect()
	if err != nil {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
		event.Error = true
		e.chaos.event <- event
		return
	}

	if inspect.State.OOMKilled == true {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "OOMKilled"
		event.Error = true
		e.chaos.event <- event
		return
	}

	if inspect.State.Dead == true {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "dead"
		event.Error = true
		e.chaos.event <- event
		return
	}

	if e.chaos.containerStopped == false && inspect.State.ExitCode != 0 {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "exit code: " + strconv.Itoa(inspect.State.ExitCode)
		event.Error = true
		e.chaos.event <- event
		return
	}

	if (e.chaos.containerStopped == true || e.chaos.containerPaused == true) != true {

		if inspect.State.Running == false {
			_, lineNumber = e.traceCodeLine()
			event.clear()
			event.ContainerName = e.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "not running"
			event.Error = true
			e.chaos.event <- event
			return
		}

		if inspect.State.Paused == true {
			_, lineNumber = e.traceCodeLine()
			event.clear()
			event.ContainerName = e.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "paused"
			event.Error = true
			e.chaos.event <- event
			return
		}

		if inspect.State.Restarting == true {
			_, lineNumber = e.traceCodeLine()
			event.clear()
			event.ContainerName = e.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "restarting"
			event.Error = true
			e.chaos.event <- event
			return
		}

	}

	logs, err = e.GetContainerLog()
	if err != nil {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
		event.Error = true
		e.chaos.event <- event
		return
	}

	lineList = e.logsCleaner(logs)
	err = e.writeContainerLogToFile(e.chaos.logPath, lineList)
	if err != nil {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
		event.Error = true
		e.chaos.event <- event
		return
	}

	line, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterFail)
	if found == true {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = string(line)
		event.Fail = true
		e.chaos.event <- event
	}

	line, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterSuccess)
	if found == true {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = string(line)
		event.Done = true
		e.chaos.event <- event
	}

	if e.chaos.enableChaos == false {
		return
	}

	if e.chaos.caosCanRestart == false {
		if e.chaos.Restart != nil && e.chaos.Restart.FilterToStart == nil && (e.chaos.Restart.TimeToStart.Min > 0 || e.chaos.Restart.TimeToStart.Max > 0) {
			e.chaos.caosCanRestart = true
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.Restart.TimeToStart.Max, e.chaos.Restart.TimeToStart.Min)
			e.chaos.Restart.minimumEventTime = time.Now().Add(timeToNextEvent)
		} else if e.chaos.Restart != nil {
			_, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.Restart.FilterToStart)
			if e.chaos.caosCanRestart == false {
				if found == true {
					e.chaos.caosCanRestart = true
				}
			}
		}
	}

	// flag o caos pode ser inicializado
	_, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterToStart)
	if e.chaos.caosStarted == false && found == true {
		e.chaos.caosStarted = true
		timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToStart, e.chaos.minimumTimeToStart)
		e.chaos.eventNext = time.Now().Add(timeToNextEvent)
	}

	if e.chaos.containerStarted == false {
		return
	}

	var restartEnable = time.Now().After(e.chaos.Restart.minimumEventTime) == true || time.Now().Equal(e.chaos.Restart.minimumEventTime) == true

	if time.Now().After(e.chaos.eventNext) == true || time.Now().Equal(e.chaos.eventNext) == true {

		if e.chaos.containerPaused == true {

			log.Printf("unpause()")
			e.chaos.containerPaused = false
			err = e.ContainerUnpause()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToPause, e.chaos.minimumTimeToPause)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		} else if e.chaos.containerStopped == true {

			log.Printf("start()")
			e.chaos.containerStopped = false
			err = e.ContainerStart()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToPause, e.chaos.minimumTimeToPause)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		} else if restartEnable == true && e.chaos.caosCanRestart == true && e.chaos.Restart != nil && e.chaos.Restart.RestartProbability >= probality && e.chaos.Restart.RestartLimit > 0 {

			log.Printf("stop()")
			e.chaos.containerStopped = true
			err = e.ContainerStop()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			e.chaos.Restart.RestartLimit -= 1
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToStart, e.chaos.minimumTimeToStart)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		} else {

			log.Printf("pause()")
			e.chaos.containerPaused = true
			err = e.ContainerPause()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToUnpause, e.chaos.minimumTimeToUnpause)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		}
	}
}

// writeContainerLogToFile
//
// Português: Escreve um arquivo csv com dados capturados da saída padrão do container e dados estatísticos do container
//   Entrada:
//     path: caminho do arquivo a ser salvo.
//     configuration: configuração do log
//       Docker: objeto padrão ContainerBuilder
//       Log: Array de LogFilter
//         Match: Texto procurado na saída padrão (tudo ou nada) de baixo para cima
//         Filter: Expressão regular contendo o filtro para isolar o texto procurado
//           Exemplo:
//             Saída padrão do container: H2021-08-20T23:46:37.586796376Z 2021/08/20 23:46:37 10.5% concluido
//             Match:   "% concluido" - Atenção: não é expressão regular
//             Filter:  "^(.*?)(?P<valueToGet>\d+)(% concluido.*)" - Atenção: Essa é uma expressão regular com nome "?P<valueToGet>"
//             Search:  "." - Numeros com pontos podem não ser bem exportados em casos como o excel, por isto, "." será substituído por ","
//             Replace: ","
//         Fail: Texto simples impresso na saída padrão indicando um erro ou bug no projeto original
//             Match:   "bug:"
//         End: Texto simples impresso na saída padrão indicando fim do teste
//             Match:   "fim!"
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
		util.TraceToLog()
		return
	}

	defer file.Close()

	var skipMatch = make([]bool, len(e.chaos.filterLog))

	var stats = types.Stats{}
	stats, err = e.ContainerStatisticsOneShot()
	if err != nil {
		util.TraceToLog()
		return
	}

	// time
	if makeLabel == true && e.logFlags&KReadingTime == KReadingTime {
		_, err = file.Write([]byte("Reading time\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KReadingTime == KReadingTime {
		_, err = file.Write([]byte(stats.Read.String()))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	for logLine := len(lineList) - 1; logLine >= 0; logLine -= 1 {
		for filterLine := 0; filterLine != len(e.chaos.filterLog); filterLine += 1 {
			if skipMatch[filterLine] == true {
				continue
			}

			if bytes.Contains(lineList[logLine], []byte(e.chaos.filterLog[filterLine].Match)) == true {
				skipMatch[filterLine] = true

				var re *regexp.Regexp
				re, err = regexp.Compile(e.chaos.filterLog[filterLine].Filter)
				if err != nil {
					util.TraceToLog()
					log.Printf("regexp.Compile().error: %v", err)
					log.Printf("regexp.Compile().error.filter: %v", e.chaos.filterLog[filterLine].Filter)
					continue
				}

				var toFile []byte
				toFile = re.ReplaceAll(lineList[logLine], []byte("${valueToGet}"))

				if e.chaos.filterLog[filterLine].Search != "" {
					re, err = regexp.Compile(e.chaos.filterLog[filterLine].Search)
					if err != nil {
						util.TraceToLog()
						log.Printf("regexp.Compile().error: %v", err)
						log.Printf("regexp.Compile().error.filter: %v", e.chaos.filterLog[filterLine].Search)
						continue
					}

					toFile = re.ReplaceAll(toFile, []byte(e.chaos.filterLog[filterLine].Replace))
				}

				if makeLabel == true {
					_, err = file.Write([]byte(e.chaos.filterLog[filterLine].Label))
					if err != nil {
						util.TraceToLog()
						return
					}

					_, err = file.Write([]byte("\t"))
					if err != nil {
						util.TraceToLog()
						return
					}
				} else {
					_, err = file.Write(toFile)
					if err != nil {
						util.TraceToLog()
						return
					}

					_, err = file.Write([]byte("\t"))
					if err != nil {
						util.TraceToLog()
						return
					}
				}
			}
		}
	}

	if makeLabel == true && e.logFlags&KBlkioIoServiceBytesRecursive == KBlkioIoServiceBytesRecursive {
		log.Printf("***************************************************************")
		log.Printf("%+v", stats.BlkioStats.IoServiceBytesRecursive)
		for i := 0; i != len(stats.BlkioStats.IoServiceBytesRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io ServiceBytes Recursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoServiceBytesRecursive == KBlkioIoServiceBytesRecursive {
		log.Printf("***************************************************************")
		for i := 0; i != len(stats.BlkioStats.IoServiceBytesRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceBytesRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceBytesRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.IoServiceBytesRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceBytesRecursive[i].Value, 10)))
		}
	}

	if makeLabel == true && e.logFlags&KBlkioIoServicedRecursive == KBlkioIoServicedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServicedRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Serviced Recursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoServicedRecursive == KBlkioIoServicedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServicedRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServicedRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServicedRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.IoServicedRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServicedRecursive[i].Value, 10)))
		}
	}

	if makeLabel == true && e.logFlags&KBlkioIoQueuedRecursive == KBlkioIoQueuedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoQueuedRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Queued Recursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoQueuedRecursive == KBlkioIoQueuedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoQueuedRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoQueuedRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoQueuedRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.IoQueuedRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoQueuedRecursive[i].Value, 10)))
		}
	}

	if makeLabel == true && e.logFlags&KBlkioIoServiceTimeRecursive == KBlkioIoServiceTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServiceTimeRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Service TimeRecursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoServiceTimeRecursive == KBlkioIoServiceTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServiceTimeRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceTimeRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceTimeRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.IoServiceTimeRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceTimeRecursive[i].Value, 10)))
		}
	}

	if makeLabel == true && e.logFlags&KBlkioIoWaitTimeRecursive == KBlkioIoWaitTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoWaitTimeRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Wait TimeRecursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoWaitTimeRecursive == KBlkioIoWaitTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoWaitTimeRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoWaitTimeRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoWaitTimeRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.IoWaitTimeRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoWaitTimeRecursive[i].Value, 10)))
		}
	}

	if makeLabel == true && e.logFlags&KBlkioIoMergedRecursive == KBlkioIoMergedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoMergedRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Merged Recursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoMergedRecursive == KBlkioIoMergedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoMergedRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoMergedRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoMergedRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.IoMergedRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoMergedRecursive[i].Value, 10)))
		}
	}

	if makeLabel == true && e.logFlags&KBlkioIoTimeRecursive == KBlkioIoTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoTimeRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Time Recursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoTimeRecursive == KBlkioIoTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoTimeRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoTimeRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoTimeRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.IoTimeRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoTimeRecursive[i].Value, 10)))
		}
	}

	if makeLabel == true && e.logFlags&KBlkioSectorsRecursive == KBlkioSectorsRecursive {
		for i := 0; i != len(stats.BlkioStats.SectorsRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Sectors Recursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioSectorsRecursive == KBlkioSectorsRecursive {
		for i := 0; i != len(stats.BlkioStats.SectorsRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.SectorsRecursive[i].Major, 10)))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.SectorsRecursive[i].Minor, 10)))
			_, err = file.Write([]byte(stats.BlkioStats.SectorsRecursive[i].Op))
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.SectorsRecursive[i].Value, 10)))
		}
	}

	// Linux specific stats, not populated on Windows.
	// Current is the number of pids in the cgroup
	if makeLabel == true && e.logFlags&KCurrentNumberOfOidsInTheCGroup == KCurrentNumberOfOidsInTheCGroup {
		_, err = file.Write([]byte("Linux specific stats, not populated on Windows. Current is the number of pids in the cgroup\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KCurrentNumberOfOidsInTheCGroup == KCurrentNumberOfOidsInTheCGroup {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PidsStats.Current)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Linux specific stats, not populated on Windows.
	// Limit is the hard limit on the number of pids in the cgroup.
	// A "Limit" of 0 means that there is no limit.
	if makeLabel == true && e.logFlags&KLimitOnTheNumberOfPidsInTheCGroup == KLimitOnTheNumberOfPidsInTheCGroup {
		_, err = file.Write([]byte("Linux specific stats, not populated on Windows. Limit is the hard limit on the number of pids in the cgroup. A \"Limit\" of 0 means that there is no limit.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KLimitOnTheNumberOfPidsInTheCGroup == KLimitOnTheNumberOfPidsInTheCGroup {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PidsStats.Limit)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Total CPU time consumed.
	// Units: nanoseconds (Linux)
	// Units: 100's of nanoseconds (Windows)
	if makeLabel == true && e.logFlags&KTotalCPUTimeConsumed == KTotalCPUTimeConsumed {
		_, err = file.Write([]byte("Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTotalCPUTimeConsumed == KTotalCPUTimeConsumed {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.CPUUsage.TotalUsage)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	if len(stats.CPUStats.CPUUsage.PercpuUsage) != 0 {
		e.logCpus = len(stats.CPUStats.CPUUsage.PercpuUsage)
	}

	if e.logFlags&KTotalCPUTimeConsumedPerCore == KTotalCPUTimeConsumedPerCore {
		// Total CPU time consumed per core (Linux). Not used on Windows.
		// Units: nanoseconds.
		if e.logCpus != 0 && len(stats.CPUStats.CPUUsage.PercpuUsage) == 0 {
			if makeLabel == true {
				for cpuNumber := 0; cpuNumber != e.logCpus; cpuNumber += 1 {
					_, err = file.Write([]byte(fmt.Sprintf("Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows. CPU: %v\t", cpuNumber)))
					if err != nil {
						util.TraceToLog()
						return
					}
				}
			} else {

				for i := 0; i != e.logCpus; i += 1 {
					_, err = file.Write([]byte{0x30})
					if err != nil {
						util.TraceToLog()
						return
					}

					_, err = file.Write([]byte("\t"))
					if err != nil {
						util.TraceToLog()
						return
					}
				}
			}
		} else if e.logCpus != 0 && len(stats.CPUStats.CPUUsage.PercpuUsage) == e.logCpus {

			for cpuNumber, cpuTime := range stats.CPUStats.CPUUsage.PercpuUsage {
				if makeLabel == true {
					_, err = file.Write([]byte(fmt.Sprintf("Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows. CPU: %v\t", cpuNumber)))
					if err != nil {
						util.TraceToLog()
						return
					}
				} else {
					_, err = file.Write([]byte(fmt.Sprintf("%v", cpuTime)))
					if err != nil {
						util.TraceToLog()
						return
					}

					_, err = file.Write([]byte("\t"))
					if err != nil {
						util.TraceToLog()
						return
					}
				}
			}
		}
	}

	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	if makeLabel == true && e.logFlags&KTimeSpentByTasksOfTheCGroupInKernelMode == KTimeSpentByTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux). Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows.Not populated for Hyper-V Containers.).\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTimeSpentByTasksOfTheCGroupInKernelMode == KTimeSpentByTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.CPUUsage.UsageInKernelmode)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Time spent by tasks of the cgroup in user mode (Linux).
	// Time spent by all container processes in user mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers
	if makeLabel == true && e.logFlags&KTimeSpentByTasksOfTheCGroupInUserMode == KTimeSpentByTasksOfTheCGroupInUserMode {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux). Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers).\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTimeSpentByTasksOfTheCGroupInUserMode == KTimeSpentByTasksOfTheCGroupInUserMode {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.CPUUsage.UsageInUsermode)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// System Usage. Linux only.
	if makeLabel == true && e.logFlags&KSystemUsage == KSystemUsage {
		_, err = file.Write([]byte("System Usage. Linux only.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KSystemUsage == KSystemUsage {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.SystemUsage)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Online CPUs. Linux only.
	if makeLabel == true && e.logFlags&KOnlineCPUs == KOnlineCPUs {
		_, err = file.Write([]byte("Online CPUs. Linux only.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KOnlineCPUs == KOnlineCPUs {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.OnlineCPUs)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWithThrottlingActive == KNumberOfPeriodsWithThrottlingActive {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods with throttling active.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWithThrottlingActive == KNumberOfPeriodsWithThrottlingActive {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.Periods)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods when the container hits its throttling limit.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerHitsItsThrottlingLimit {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.ThrottledPeriods)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if makeLabel == true && e.logFlags&KAggregateTimeTheContainerWasThrottledForInNanoseconds == KAggregateTimeTheContainerWasThrottledForInNanoseconds {
		_, err = file.Write([]byte("Throttling Data. Linux only. Aggregate time the container was throttled for in nanoseconds.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KAggregateTimeTheContainerWasThrottledForInNanoseconds == KAggregateTimeTheContainerWasThrottledForInNanoseconds {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.ThrottlingData.ThrottledTime)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// CPU Usage. Linux and Windows.
	// Total CPU time consumed.
	// Units: nanoseconds (Linux)
	// Units: 100's of nanoseconds (Windows)
	if makeLabel == true && e.logFlags&KTotalPreCPUTimeConsumed == KTotalPreCPUTimeConsumed {
		_, err = file.Write([]byte("Total CPU time consumed. (Units: nanoseconds on Linux. Units: 100's of nanoseconds on Windows)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTotalPreCPUTimeConsumed == KTotalPreCPUTimeConsumed {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.CPUUsage.TotalUsage)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	if makeLabel == true && e.logFlags&KTotalPreCPUTimeConsumedPerCore == KTotalPreCPUTimeConsumedPerCore {
		for cpuNumber := 0; cpuNumber != e.logCpus; cpuNumber += 1 {
			_, err = file.Write([]byte(fmt.Sprintf("Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows. CPU: %v\t", cpuNumber)))
			if err != nil {
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
					util.TraceToLog()
					return
				}

				_, err = file.Write([]byte("\t"))
				if err != nil {
					util.TraceToLog()
					return
				}
			}
		}

		for _, cpuTime := range stats.PreCPUStats.CPUUsage.PercpuUsage {
			_, err = file.Write([]byte(fmt.Sprintf("%v", cpuTime)))
			if err != nil {
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte("\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	}

	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	if makeLabel == true && e.logFlags&KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode == KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux) - Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows - Not populated for Hyper-V Containers.)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode == KTimeSpentByPreCPUTasksOfTheCGroupInKernelMode {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.CPUUsage.UsageInKernelmode)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in user mode (Linux).
	// Time spent by all container processes in user mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers
	if makeLabel == true && e.logFlags&KTimeSpentByPreCPUTasksOfTheCGroupInUserMode == KTimeSpentByPreCPUTasksOfTheCGroupInUserMode {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux) - Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KTimeSpentByPreCPUTasksOfTheCGroupInUserMode == KTimeSpentByPreCPUTasksOfTheCGroupInUserMode {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.CPUUsage.UsageInUsermode)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// System Usage. Linux only.
	if makeLabel == true && e.logFlags&KPreCPUSystemUsage == KPreCPUSystemUsage {
		_, err = file.Write([]byte("System Usage. (Linux only)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KPreCPUSystemUsage == KPreCPUSystemUsage {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.SystemUsage)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Online CPUs. Linux only.
	if makeLabel == true && e.logFlags&KOnlinePreCPUs == KOnlinePreCPUs {
		_, err = file.Write([]byte("Online CPUs. (Linux only)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KOnlinePreCPUs == KOnlinePreCPUs {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.OnlineCPUs)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
	if makeLabel == true && e.logFlags&KAggregatePreCPUTimeTheContainerWasThrottled == KAggregatePreCPUTimeTheContainerWasThrottled {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Aggregate time the container was throttled for in nanoseconds.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KAggregatePreCPUTimeTheContainerWasThrottled == KAggregatePreCPUTimeTheContainerWasThrottled {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.ThrottledTime)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Throttling Data. Linux only.
	// Number of periods with throttling active
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWithPreCPUThrottlingActive == KNumberOfPeriodsWithPreCPUThrottlingActive {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Number of periods with throttling active.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWithPreCPUThrottlingActive == KNumberOfPeriodsWithPreCPUThrottlingActive {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.Periods)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
	if makeLabel == true && e.logFlags&KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Number of periods when the container hits its throttling limit.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit == KNumberOfPeriodsWhenTheContainerPreCPUHitsItsThrottlingLimit {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData.ThrottledPeriods)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// current res_counter usage for memory
	if makeLabel == true && e.logFlags&KCurrentResCounterUsageForMemory == KCurrentResCounterUsageForMemory {
		_, err = file.Write([]byte("Current res_counter usage for memory\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KCurrentResCounterUsageForMemory == KCurrentResCounterUsageForMemory {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Usage)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// maximum usage ever recorded.
	if makeLabel == true && e.logFlags&KMaximumUsageEverRecorded == KMaximumUsageEverRecorded {
		_, err = file.Write([]byte("Maximum usage ever recorded.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KMaximumUsageEverRecorded == KMaximumUsageEverRecorded {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.MaxUsage)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// number of times memory usage hits limits.
	if makeLabel == true && e.logFlags&KNumberOfTimesMemoryUsageHitsLimits == KNumberOfTimesMemoryUsageHitsLimits {
		_, err = file.Write([]byte("Number of times memory usage hits limits.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KNumberOfTimesMemoryUsageHitsLimits == KNumberOfTimesMemoryUsageHitsLimits {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Failcnt)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	if makeLabel == true && e.logFlags&KMemoryLimit == KMemoryLimit {
		_, err = file.Write([]byte("Memory limit\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KMemoryLimit == KMemoryLimit {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Limit)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// committed bytes
	if makeLabel == true && e.logFlags&KCommittedBytes == KCommittedBytes {
		_, err = file.Write([]byte("Committed bytes\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KCommittedBytes == KCommittedBytes {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.Commit)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// peak committed bytes
	if makeLabel == true && e.logFlags&KPeakCommittedBytes == KPeakCommittedBytes {
		_, err = file.Write([]byte("Peak committed bytes\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KPeakCommittedBytes == KPeakCommittedBytes {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.CommitPeak)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	// private working set
	if makeLabel == true && e.logFlags&KPrivateWorkingSet == KPrivateWorkingSet {
		_, err = file.Write([]byte("Private working set\n"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else if e.logFlags&KPrivateWorkingSet == KPrivateWorkingSet {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.PrivateWorkingSet)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	_, err = file.Write([]byte("\n"))
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
