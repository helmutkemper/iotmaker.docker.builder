// package theater
//
// Português: Monta um teatro de containers com senas lineares ou de caos, permitindo ao desenvolvedor montar em sua
// máquina de trabalho testes de micro caos durante o desenvolvimento das suas aplicações.
//
// Para entender melhor o uso de senários de micro caos, imagine o desenvolvimento de aplicações se comunicando com
// outras aplicações, onde a pausa e reinício aleatórios dos containers ajuda a entender melhor o
// comportamento do projeto quando algo dá errado e a comunicação falha.
package theater

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	dockerBuild "github.com/helmutkemper/iotmaker.docker.builder"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"github.com/helmutkemper/util"
	"io/fs"
	"log"
	"math/rand"
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

type Theater struct {
	imageExpirationTime time.Duration
	sceneCache          []*Configuration
	sceneBuilding       []*Configuration
	scenePrologue       []*Configuration
	ticker              *time.Ticker
	event               chan Event
	cpus                int
	logFlags            int64
}

type Timers struct {
	Min time.Duration
	Max time.Duration
}

type Event struct {
	ContainerName string
	Message       string
	Error         bool
	Done          bool
	Fail          bool
}

func (e *Event) clear() {
	e.ContainerName = ""
	e.Message = ""
	e.Done = false
	e.Error = false
	e.Fail = false
}

type LogFilter struct {
	Label string

	// Texto contido na linha (tudo ou nada)
	Match string

	// expressão regular contendo o filtro para capturar o elemento
	// Ex.: ^(.*?)(?P<valueToGet>\\d+)(.*)
	Filter string

	// texto usado em replaceAll
	// Ex.: search: "." replace: "," para compatibilizar número com o excel
	Search  string
	Replace string
}

type Restart struct {
	FilterToStart      []LogFilter
	TimeToStart        Timers
	RestartProbability float64
	RestartLimit       int

	minimumEventTime time.Time
}

type Chaos struct {
	FilterToStart []LogFilter
	Restart       *Restart
	TimeToStart   Timers
	TimeToPause   Timers
	TimeToUnpause Timers
}

type Configuration struct {
	Docker  *dockerBuild.ContainerBuilder
	LogPath string
	Log     []LogFilter
	Fail    []LogFilter
	End     []LogFilter

	Chaos Chaos

	linear            bool
	caosStarted       bool
	caosCanRestart    bool
	caosCanRestartEnd bool

	containerStarted bool
	containerPaused  bool
	containerStopped bool

	eventNext time.Time
}

// NewTestContainerConfiguration
//
// Português: Adiciona um objeto de configuração de container para ser usado no teatro de teste.
//   Entrada:
//     docker: Ponteiro para o objeto dockerBuild.ContainerBuilder com a configuração do container a ser gerado.
//   Saída:
//     configuration: Objeto de configuração preenchido.
func NewTestContainerConfiguration(docker *dockerBuild.ContainerBuilder) (configuration *Configuration) {
	return &Configuration{Docker: docker}
}

// SetASceneLinearFlag
//
// Português: Transforma o senário de teste em um teste linear, preparando e subindo o container.
func (e *Configuration) SetASceneLinearFlag() (configuration *Configuration) {
	e.linear = true
	return e
}

// SetContainerStatsLogPath
//
// Português: Salva um arquivo CSV com dados de métricas do container.
//   Entrada:
//     path: Caminho de onde salvar o arquivo CSV
//   Saída:
//     configuration: Objeto de configuração preenchido.
func (e *Configuration) SetContainerStatsLogPath(path string) (configuration *Configuration) {
	e.LogPath = path
	return e
}

// AddFilterToCaptureInformationOnTheContainersStandardOutputForStatsLog
//
// Adiciona um filtro de busca na saída padrão do container e arquiva a informação no arquivo CSV com as medições de uso
// do container.
//
//   Entrada:
//     label: Rótudo adicionado ao arquivo CSV.
//     match: Texto simples procurado saída padrão a fim de indicar ocorência.
//     filter: Expressão regular nomeada, com o termo `valueToGet`, usada para isolar o valor a ser escrito no arquivo CSV.
//     search: Rexpressão regular usada para substituir valores no resultado de `filter` [opcional]
//     replace: Elemento replace da expressão regular usada em `search` [opcional]
//   Saída:
//     configuration: Objeto de configuração preenchido.
//
// Nota: - Esta configuração requer a configuração SetContainerStatsLogPath()
func (e *Configuration) AddFilterToCaptureInformationOnTheContainersStandardOutputForStatsLog(label, match, filter, search, replace string) (configuration *Configuration) {

	if e.Log == nil {
		e.Log = make([]LogFilter, 0)
	}

	e.Log = append(
		e.Log, LogFilter{
			Label:   label,
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

func (e *Configuration) AddASceneSettingFilterOnTheContainersStandardOutputToFailEvent(match, filter, search, replace string) (configuration *Configuration) {

	if e.Fail == nil {
		e.Fail = make([]LogFilter, 0)
	}

	e.Fail = append(
		e.Fail, LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

func (e *Configuration) AddASceneSettingFilterOnTheContainersStandardOutputToEndEvent(match, filter, search, replace string) (configuration *Configuration) {

	if e.End == nil {
		e.End = make([]LogFilter, 0)
	}

	e.End = append(
		e.End, LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

func (e *Configuration) AddASceneOfChaosSettingFilterOnTheContainersStandardOutputToStartCaos(match, filter, search, replace string) (configuration *Configuration) {

	if e.Chaos.FilterToStart == nil {
		e.Chaos.FilterToStart = make([]LogFilter, 0)
	}

	e.Chaos.FilterToStart = append(
		e.Chaos.FilterToStart, LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

// AddASceneOfChaosSettingPauseDuration
func (e *Configuration) AddASceneOfChaosSettingPauseDuration(min, max time.Duration) (configuration *Configuration) {

	e.Chaos.TimeToPause.Min = min
	e.Chaos.TimeToPause.Max = max

	if e.Chaos.TimeToStart.Min == 0 && e.Chaos.TimeToStart.Max == 0 {
		e.Chaos.TimeToStart.Min = min
		e.Chaos.TimeToStart.Max = max
	}

	if e.Chaos.TimeToUnpause.Min == 0 && e.Chaos.TimeToUnpause.Max == 0 {
		e.Chaos.TimeToUnpause.Min = min
		e.Chaos.TimeToUnpause.Max = max
	}

	return e
}

func (e *Configuration) AddASceneOfChaosSettingUnpauseDuration(min, max time.Duration) (configuration *Configuration) {

	e.Chaos.TimeToUnpause.Min = min
	e.Chaos.TimeToUnpause.Max = max

	if e.Chaos.TimeToStart.Min == 0 && e.Chaos.TimeToStart.Max == 0 {
		e.Chaos.TimeToStart.Min = min
		e.Chaos.TimeToStart.Max = max
	}

	if e.Chaos.TimeToPause.Min == 0 && e.Chaos.TimeToPause.Max == 0 {
		e.Chaos.TimeToPause.Min = min
		e.Chaos.TimeToPause.Max = max
	}

	return e
}

func (e *Configuration) AddASceneOfChaosSettingStartDuration(min, max time.Duration) (configuration *Configuration) {

	e.Chaos.TimeToStart.Min = min
	e.Chaos.TimeToStart.Max = max

	if e.Chaos.TimeToUnpause.Min == 0 && e.Chaos.TimeToUnpause.Max == 0 {
		e.Chaos.TimeToUnpause.Min = min
		e.Chaos.TimeToUnpause.Max = max
	}

	if e.Chaos.TimeToPause.Min == 0 && e.Chaos.TimeToPause.Max == 0 {
		e.Chaos.TimeToPause.Min = min
		e.Chaos.TimeToPause.Max = max
	}

	return e
}

func (e *Configuration) AddASceneOfChaosSettingStartDurationRestartFilterOnTheContainersStandardOutput(match, filter, search, replace string) (configuration *Configuration) {

	if e.Chaos.Restart == nil {
		e.Chaos.Restart = &Restart{}
	}

	if e.Chaos.Restart.FilterToStart == nil {
		e.Chaos.Restart.FilterToStart = make([]LogFilter, 0)
	}

	e.Chaos.Restart.FilterToStart = append(
		e.Chaos.Restart.FilterToStart, LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

func (e *Configuration) AddASceneOfChaosSettingRestartInterval(min, max time.Duration) (configuration *Configuration) {

	if e.Chaos.Restart == nil {
		e.Chaos.Restart = &Restart{}
	}

	e.Chaos.Restart.TimeToStart.Min = min
	e.Chaos.Restart.TimeToStart.Max = max

	return e
}

func (e *Configuration) AddASceneOfChaosSettingRestartIntervalRestartController(probability float64, limit int) (configuration *Configuration) {

	if e.Chaos.Restart == nil {
		e.Chaos.Restart = &Restart{}
	}

	e.Chaos.Restart.RestartProbability = probability
	e.Chaos.Restart.RestartLimit = limit

	return e
}

func (e *Theater) SetImageExpirationTime(expiration time.Duration) {
	e.imageExpirationTime = expiration
}

func (e *Theater) GetImageExpirationTime() (expiration time.Duration) {
	return e.imageExpirationTime
}

func (e *Theater) imageExpirationTimeIsValid(docker *dockerBuild.ContainerBuilder) (valid bool) {
	return time.Now().Add(e.GetImageExpirationTime() * -1).Before(docker.GetImageCreated())
}

func (e *Theater) GetChannels() (eventChannel *chan Event) {
	return &e.event
}

func (e *Theater) Init() (err error) {

	if e.sceneBuilding == nil {
		err = errors.New("reference of scenes is empty")
		return
	}

	var osName = runtime.GOOS
	if e.logFlags == 0 && osName == "darwin" {
		e.logFlags = KMacOsLog
	} else if e.logFlags == 0 {
		e.logFlags = KAll
	}

	e.event = make(chan Event)

	err = e.buildAll()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.startPrologueScene()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.startCaosScene()
	if err != nil {
		util.TraceToLog()
		return
	}

	e.ticker = time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-e.ticker.C:
				e.manager()
			}
		}
	}()

	return
}

func (e *Theater) logsCleaner(logs []byte) [][]byte {
	logs = bytes.ReplaceAll(logs, []byte("\r"), []byte(""))
	return bytes.Split(logs, []byte("\n"))
}

func (e *Theater) logsSearchAndReplaceIntoText(lineList [][]byte, configuration []LogFilter) (line []byte, found bool) {
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

func (e *Theater) manager() {
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

	for _, container := range e.sceneBuilding {

		probality = e.getProbalityNumber()

		inspect, err = container.Docker.ContainerInspect()
		if err != nil {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
			event.Error = true
			e.event <- event
			continue
		}

		if inspect.State.OOMKilled == true {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "OOMKilled"
			event.Error = true
			e.event <- event
			continue
		}

		if inspect.State.Dead == true {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "dead"
			event.Error = true
			e.event <- event
			continue
		}

		if container.containerStopped == false && inspect.State.ExitCode != 0 {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "exit code: " + strconv.Itoa(inspect.State.ExitCode)
			event.Error = true
			e.event <- event
			continue
		}

		if (container.containerStopped == true || container.containerPaused == true) != true {

			if inspect.State.Running == false {
				lineNumber = traceLineFromCode()
				event.clear()
				event.ContainerName = container.Docker.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "not running"
				event.Error = true
				e.event <- event
				continue
			}

			if inspect.State.Paused == true {
				lineNumber = traceLineFromCode()
				event.clear()
				event.ContainerName = container.Docker.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "paused"
				event.Error = true
				e.event <- event
				continue
			}

			if inspect.State.Restarting == true {
				lineNumber = traceLineFromCode()
				event.clear()
				event.ContainerName = container.Docker.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + "restarting"
				event.Error = true
				e.event <- event
				continue
			}

		}

		logs, err = container.Docker.GetContainerLog()
		if err != nil {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
			event.Error = true
			e.event <- event
			continue
		}

		lineList = e.logsCleaner(logs)
		err = e.writeContainerLogToFile(container.LogPath, lineList, container)
		if err != nil {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
			event.Error = true
			e.event <- event
			continue
		}

		line, found = e.logsSearchAndReplaceIntoText(lineList, container.Fail)
		if found == true {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = string(line)
			event.Fail = true
			e.event <- event
		}

		line, found = e.logsSearchAndReplaceIntoText(lineList, container.End)
		if found == true {
			lineNumber = traceLineFromCode()
			event.clear()
			event.ContainerName = container.Docker.GetContainerName()
			event.Message = string(line)
			event.Done = true
			e.event <- event
		}

		if container.linear == true {
			continue
		}

		if container.caosCanRestart == false {
			if container.Chaos.Restart != nil && container.Chaos.Restart.FilterToStart == nil && (container.Chaos.Restart.TimeToStart.Min > 0 || container.Chaos.Restart.TimeToStart.Max > 0) {
				container.caosCanRestart = true
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.Restart.TimeToStart.Max, container.Chaos.Restart.TimeToStart.Min)
				container.Chaos.Restart.minimumEventTime = time.Now().Add(timeToNextEvent)
			} else if container.Chaos.Restart != nil {
				_, found = e.logsSearchAndReplaceIntoText(lineList, container.Chaos.Restart.FilterToStart)
				if container.caosCanRestart == false {
					if found == true {
						container.caosCanRestart = true
					}
				}
			}
		}

		// flag o caos pode ser inicializado
		_, found = e.logsSearchAndReplaceIntoText(lineList, container.Chaos.FilterToStart)
		if container.caosStarted == false && found == true {
			container.caosStarted = true
			timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToStart.Max, container.Chaos.TimeToStart.Min)
			container.eventNext = time.Now().Add(timeToNextEvent)
		}

		if container.containerStarted == false {
			continue
		}

		var restartEnable = time.Now().After(container.Chaos.Restart.minimumEventTime) == true || time.Now().Equal(container.Chaos.Restart.minimumEventTime) == true

		if time.Now().After(container.eventNext) == true || time.Now().Equal(container.eventNext) == true {

			if container.containerPaused == true {

				log.Printf("unpause()")
				container.containerPaused = false
				err = container.Docker.ContainerUnpause()
				if err != nil {
					lineNumber = traceLineFromCode()
					event.clear()
					event.ContainerName = container.Docker.GetContainerName()
					event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
					event.Error = true
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToPause.Max, container.Chaos.TimeToPause.Min)
				container.eventNext = time.Now().Add(timeToNextEvent)

			} else if container.containerStopped == true {

				log.Printf("start()")
				container.containerStopped = false
				err = container.Docker.ContainerStart()
				if err != nil {
					lineNumber = traceLineFromCode()
					event.clear()
					event.ContainerName = container.Docker.GetContainerName()
					event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
					event.Error = true
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToPause.Max, container.Chaos.TimeToPause.Min)
				container.eventNext = time.Now().Add(timeToNextEvent)

			} else if restartEnable == true && container.caosCanRestart == true && container.Chaos.Restart != nil && container.Chaos.Restart.RestartProbability >= probality && container.Chaos.Restart.RestartLimit > 0 {

				log.Printf("stop()")
				container.containerStopped = true
				err = container.Docker.ContainerStop()
				if err != nil {
					lineNumber = traceLineFromCode()
					event.clear()
					event.ContainerName = container.Docker.GetContainerName()
					event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
					event.Error = true
					continue
				}
				container.Chaos.Restart.RestartLimit -= 1
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToStart.Max, container.Chaos.TimeToStart.Min)
				container.eventNext = time.Now().Add(timeToNextEvent)

			} else {

				log.Printf("pause()")
				container.containerPaused = true
				err = container.Docker.ContainerPause()
				if err != nil {
					lineNumber = traceLineFromCode()
					event.clear()
					event.ContainerName = container.Docker.GetContainerName()
					event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
					event.Error = true
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToUnpause.Max, container.Chaos.TimeToUnpause.Min)
				container.eventNext = time.Now().Add(timeToNextEvent)

			}
		}
	}
}

func (e *Theater) AddCacheConfig(container *Configuration) (err error) {
	if e.sceneCache == nil {
		e.sceneCache = make([]*Configuration, 0)
	}

	if container.Docker.GetInitialized() == true {
		util.TraceToLog()
		err = errors.New("do not initialize the container")
		return
	}

	var imageName = container.Docker.GetImageName()
	if imageName == "" {
		util.TraceToLog()
		err = errors.New("image name must be set")
		return
	}

	var buildPath = container.Docker.GetBuildFolderPath()
	var url = container.Docker.GetGitCloneToBuild()

	if buildPath == "" && url == "" {
		util.TraceToLog()
		err = errors.New("image folder path or git server path must be set")
		return
	}

	if buildPath != "" && url != "" {
		util.TraceToLog()
		err = errors.New("select once one, image folder path or git server path")
		return
	}

	e.sceneCache = append(e.sceneCache, container)

	return
}

func (e *Theater) buildCache() (err error) {
	var id string

	for _, container := range e.sceneCache {
		id, err = container.Docker.ImageFindIdByName(container.Docker.GetImageCacheName())
		if err != nil {
			util.TraceToLog()
			return
		}

		if id != "" {
			continue
		}

		err = container.Docker.Init()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.buildImage(container.Docker)
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *Theater) AddContainerConfiguration(container *Configuration) (err error) {
	if e.sceneBuilding == nil {
		e.sceneBuilding = make([]*Configuration, 0)
	}

	if container.Docker.GetInitialized() == true {
		err = errors.New("do not initialize the container")
		return
	}

	var imageName = container.Docker.GetImageName()
	if imageName == "" {
		err = errors.New("image name must be set")
		return
	}

	var containerName = container.Docker.GetContainerName()
	if containerName == "" {
		err = errors.New("container name must be set")
		return
	}

	e.sceneBuilding = append(e.sceneBuilding, container)

	return
}

func (e *Theater) buildImage(docker *dockerBuild.ContainerBuilder) (err error) {
	var buildPath = docker.GetBuildFolderPath()
	var url = docker.GetGitCloneToBuild()

	if buildPath != "" && (e.GetImageExpirationTime() == 0 || e.imageExpirationTimeIsValid(docker) == false) {
		_, err = docker.ImageBuildFromFolder()
	} else if url != "" && (e.GetImageExpirationTime() == 0 || e.imageExpirationTimeIsValid(docker) == false) {
		_, err = docker.ImageBuildFromServer()
	}
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

func (e *Theater) buildContainers() (err error) {
	//var id string

	for _, container := range e.sceneBuilding {
		//id, err = container.Docker.ImageFindIdByName(container.Docker.GetImageName())
		//if err != nil && err.Error() != "image name not found" {
		//	util.TraceToLog()
		//	return
		//}

		err = container.Docker.Init()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.buildImage(container.Docker)
		if err != nil {
			util.TraceToLog()
			return
		}

		err = container.Docker.ContainerBuildWithoutStartingItFromImage()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *Theater) AddContainerConfigurationToPrologueScene(container *Configuration) (err error) {
	if e.scenePrologue == nil {
		e.scenePrologue = make([]*Configuration, 0)
	}

	if container.Docker.GetInitialized() == true {
		util.TraceToLog()
		err = errors.New("do not initialize the container")
		return
	}

	e.scenePrologue = append(e.scenePrologue, container)

	return
}

func (e *Theater) startPrologueScene() (err error) {
	for _, container := range e.scenePrologue {
		err = container.Docker.ContainerStartAfterBuild()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *Theater) startCaosScene() (err error) {

	if e.sceneBuilding == nil {
		err = errors.New("reference of scenes is empty")
		return
	}

	for _, container := range e.sceneBuilding {

		if container.Docker.GetInitialized() == false {
			err = container.Docker.Init()
			if err != nil {
				util.TraceToLog()
				return
			}
		}

		container.containerStarted = true

		if container.Docker.GetContainerIsStarted() == false {
			err = container.Docker.ContainerStartAfterBuild()
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	}

	return
}

func (e *Theater) buildAll() (err error) {
	if e.sceneBuilding == nil {
		err = errors.New("reference of scenes is empty")
		return
	}

	err = e.buildCache()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.buildContainers()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
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
func (e *Theater) writeContainerLogToFile(path string, lineList [][]byte, configuration *Configuration) (err error) {
	if path == "" {
		return
	}

	if lineList == nil {
		return
	}

	if configuration == nil {
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

	var skipMatch = make([]bool, len(configuration.Log))

	var stats = types.Stats{}
	stats, err = configuration.Docker.ContainerStatisticsOneShot()
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
		for filterLine := 0; filterLine != len(configuration.Log); filterLine += 1 {
			if skipMatch[filterLine] == true {
				continue
			}

			if bytes.Contains(lineList[logLine], []byte(configuration.Log[filterLine].Match)) == true {
				skipMatch[filterLine] = true

				var re *regexp.Regexp
				re, err = regexp.Compile(configuration.Log[filterLine].Filter)
				if err != nil {
					util.TraceToLog()
					log.Printf("regexp.Compile().error: %v", err)
					log.Printf("regexp.Compile().error.filter: %v", configuration.Log[filterLine].Filter)
					continue
				}

				var toFile []byte
				toFile = re.ReplaceAll(lineList[logLine], []byte("${valueToGet}"))

				if configuration.Log[filterLine].Search != "" {
					re, err = regexp.Compile(configuration.Log[filterLine].Search)
					if err != nil {
						util.TraceToLog()
						log.Printf("regexp.Compile().error: %v", err)
						log.Printf("regexp.Compile().error.filter: %v", configuration.Log[filterLine].Search)
						continue
					}

					toFile = re.ReplaceAll(toFile, []byte(configuration.Log[filterLine].Replace))
				}

				if makeLabel == true {
					_, err = file.Write([]byte(configuration.Log[filterLine].Label))
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
		e.cpus = len(stats.CPUStats.CPUUsage.PercpuUsage)
	}

	if e.logFlags&KTotalCPUTimeConsumedPerCore == KTotalCPUTimeConsumedPerCore {
		// Total CPU time consumed per core (Linux). Not used on Windows.
		// Units: nanoseconds.
		if e.cpus != 0 && len(stats.CPUStats.CPUUsage.PercpuUsage) == 0 {
			if makeLabel == true {
				for cpuNumber := 0; cpuNumber != e.cpus; cpuNumber += 1 {
					_, err = file.Write([]byte(fmt.Sprintf("Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows. CPU: %v\t", cpuNumber)))
					if err != nil {
						util.TraceToLog()
						return
					}
				}
			} else {

				for i := 0; i != e.cpus; i += 1 {
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
		} else if e.cpus != 0 && len(stats.CPUStats.CPUUsage.PercpuUsage) == e.cpus {

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
		for cpuNumber := 0; cpuNumber != e.cpus; cpuNumber += 1 {
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
		if e.cpus != 0 && len(stats.CPUStats.CPUUsage.PercpuUsage) == 0 {
			for i := 0; i != e.cpus; i += 1 {
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

func (e *Theater) getProbalityNumber() (probality float64) {
	return 1.0 - e.getRandSeed().Float64()
}

func (e *Theater) SetLogFields(logFlags int64) {
	e.logFlags = logFlags
}

func (e *Theater) selectBetweenMaxAndMin(max, min time.Duration) (selected time.Duration) {
	randValue := e.getRandSeed().Int63n(int64(max)-int64(min)) + int64(min)
	return time.Duration(randValue)
}

func (e *Theater) getRandSeed() (seed *rand.Rand) {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

func traceLineFromCode() (line int) {
	var ok bool

	_, _, line, ok = runtime.Caller(1)
	if !ok {
		log.Printf("TraceToLog().error: runtime.Caller() unknown error")
		return
	}

	return
}
