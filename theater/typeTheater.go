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

type Theater struct {
	sceneCache    []*Configuration
	sceneBuilding []*Configuration
	scenePrologue []*Configuration
	sceneCaos     []*Configuration

	refCaos []*Configuration

	ticker *time.Ticker

	errchannel chan error

	cpus int
}

type Timers struct {
	Min time.Duration
	Max time.Duration
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

	Linear            bool
	caosStarted       bool
	caosCanRestart    bool
	caosCanRestartEnd bool

	containerStarted bool
	containerPaused  bool
	containerStopped bool

	testEnd bool

	eventNext time.Time
}

func NewTestContainerConfiguration(docker *dockerBuild.ContainerBuilder) (configuration *Configuration) {
	return &Configuration{Docker: docker}
}

func (e *Configuration) SetASceneLinearFlag() (configuration *Configuration) {
	e.Linear = true
	return e
}

func (e *Configuration) SetContainerStatsLogPath(path string) (configuration *Configuration) {
	e.LogPath = path
	return e
}

//Add a filter to capture information on the container's standard output for stats log
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

func (e *Configuration) AddASineOfChaosSettingFilterOnTheContainersStandardOutputToFailEvent(match, filter, search, replace string) (configuration *Configuration) {

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

func (e *Configuration) AddASineOfChaosSettingFilterOnTheContainersStandardOutputToEndEvent(match, filter, search, replace string) (configuration *Configuration) {

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

func (e *Configuration) AddASineOfChaosSettingFilterOnTheContainersStandardOutputToStartCaos(match, filter, search, replace string) (configuration *Configuration) {

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

// AddASineOfChaosSettingPauseDuration
func (e *Configuration) AddASineOfChaosSettingPauseDuration(min, max time.Duration) (configuration *Configuration) {

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

func (e *Configuration) AddASineOfChaosSettingUnpauseDuration(min, max time.Duration) (configuration *Configuration) {

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

func (e *Configuration) AddASineOfChaosSettingStartDuration(min, max time.Duration) (configuration *Configuration) {

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

func (e *Configuration) AddASineOfChaosSettingStartDurationRestartFilterOnTheContainersStandardOutput(match, filter, search, replace string) (configuration *Configuration) {

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

func (e *Configuration) AddASineOfChaosSettingRestartInterval(min, max time.Duration) (configuration *Configuration) {

	if e.Chaos.Restart == nil {
		e.Chaos.Restart = &Restart{}
	}

	e.Chaos.Restart.TimeToStart.Min = min
	e.Chaos.Restart.TimeToStart.Max = max

	return e
}

func (e *Configuration) AddASineOfChaosSettingRestartIntervalRestartController(probability float64, limit int) (configuration *Configuration) {

	if e.Chaos.Restart == nil {
		e.Chaos.Restart = &Restart{}
	}

	e.Chaos.Restart.RestartProbability = probability
	e.Chaos.Restart.RestartLimit = limit

	return e
}

func (e *Theater) Init() (err error) {

	e.errchannel = make(chan error)

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

	go func() {
		for {
			select {
			case err := <-e.errchannel:
				log.Printf("error: %v", err)
			}
		}
	}()

	return
}

func (e *Theater) logsCleaner(logs []byte) [][]byte {
	logs = bytes.ReplaceAll(logs, []byte("\r"), []byte(""))
	return bytes.Split(logs, []byte("\n"))
}

func (e *Theater) logsSearchSimplesText(lineList [][]byte, configuration []LogFilter) (line []byte, found bool) {
	if configuration == nil {
		return
	}

	for logLine := len(lineList) - 1; logLine >= 0; logLine -= 1 {

		for filterLine := 0; filterLine != len(configuration); filterLine += 1 {
			line = lineList[logLine]
			if bytes.Contains(line, []byte(configuration[filterLine].Match)) == true {
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

	var inspect iotmakerdocker.ContainerInspect

	for _, container := range e.refCaos {

		probality = e.getProbalityNumber()

		inspect, err = container.Docker.ContainerInspect()
		if err != nil {
			lineNumber = traceLineFromCode()
			e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: " + err.Error())
			continue
		}

		if inspect.State.OOMKilled == true {
			lineNumber = traceLineFromCode()
			e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: OOMKilled")
			continue
		}

		if inspect.State.Dead == true {
			lineNumber = traceLineFromCode()
			e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: dead")
			continue
		}

		if container.containerStopped == false && inspect.State.ExitCode != 0 {
			lineNumber = traceLineFromCode()
			e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: exit code " + strconv.Itoa(inspect.State.ExitCode))
			continue
		}

		if (container.containerStopped == true || container.containerPaused == true) != true {

			if inspect.State.Running == false {
				lineNumber = traceLineFromCode()
				e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: not running")
				continue
			}

			if inspect.State.Paused == true {
				lineNumber = traceLineFromCode()
				e.errchannel <- errors.New(container.Docker.GetContainerName() + ".error: paused")
				continue
			}

			if inspect.State.Restarting == true {
				lineNumber = traceLineFromCode()
				e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: restarting")
				continue
			}

		}

		logs, err = container.Docker.GetContainerLog()

		if err != nil {
			lineNumber = traceLineFromCode()
			e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: " + err.Error())
			continue
		}

		lineList = e.logsCleaner(logs)

		err = e.writeContainerLogToFile(container.LogPath, lineList, container)
		if err != nil {
			lineNumber = traceLineFromCode()
			e.errchannel <- errors.New(strconv.Itoa(lineNumber) + " - " + container.Docker.GetContainerName() + ".error: " + err.Error())
			continue
		}

		if container.Linear == true {
			continue
		}

		if container.caosCanRestart == false {
			if container.Chaos.Restart != nil && container.Chaos.Restart.FilterToStart == nil && (container.Chaos.Restart.TimeToStart.Min > 0 || container.Chaos.Restart.TimeToStart.Max > 0) {
				container.caosCanRestart = true
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.Restart.TimeToStart.Max, container.Chaos.Restart.TimeToStart.Min)
				container.Chaos.Restart.minimumEventTime = time.Now().Add(timeToNextEvent)
			} else if container.Chaos.Restart != nil {
				_, found = e.logsSearchSimplesText(lineList, container.Chaos.Restart.FilterToStart)
				if container.caosCanRestart == false {
					if found == true {
						container.caosCanRestart = true
					}
				}
			}
		}

		// flag o caos pode ser inicializado
		_, found = e.logsSearchSimplesText(lineList, container.Chaos.FilterToStart)
		if container.caosStarted == false && found == true {
			container.caosStarted = true
			timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToStart.Max, container.Chaos.TimeToStart.Min)
			container.eventNext = time.Now().Add(timeToNextEvent)
		}

		if container.containerStarted == false {
			continue
		}

		line, found = e.logsSearchSimplesText(lineList, container.Fail)
		if found == true {
			e.errchannel <- errors.New(container.Docker.GetContainerName() + ".error: test fail - " + string(line))
		}

		line, found = e.logsSearchSimplesText(lineList, container.End)
		if found == true {
			container.testEnd = true
		}

		var restartEnable = time.Now().After(container.Chaos.Restart.minimumEventTime) == true || time.Now().Equal(container.Chaos.Restart.minimumEventTime) == true

		if time.Now().After(container.eventNext) == true || time.Now().Equal(container.eventNext) == true {

			if container.containerPaused == true {

				log.Printf("unpause()")
				container.containerPaused = false
				err = container.Docker.ContainerUnpause()
				if err != nil {
					e.errchannel <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToPause.Max, container.Chaos.TimeToPause.Min)
				container.eventNext = time.Now().Add(timeToNextEvent)

			} else if container.containerStopped == true {

				log.Printf("start()")
				container.containerStopped = false
				err = container.Docker.ContainerStart()
				if err != nil {
					e.errchannel <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					util.TraceToLog()
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToPause.Max, container.Chaos.TimeToPause.Min)
				container.eventNext = time.Now().Add(timeToNextEvent)

			} else if restartEnable == true && container.caosCanRestart == true && container.Chaos.Restart != nil && container.Chaos.Restart.RestartProbability >= probality && container.Chaos.Restart.RestartLimit > 0 {

				log.Printf("stop()")
				container.containerStopped = true
				err = container.Docker.ContainerStop()
				if err != nil {
					e.errchannel <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					util.TraceToLog()
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
					e.errchannel <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					util.TraceToLog()
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Chaos.TimeToUnpause.Max, container.Chaos.TimeToUnpause.Min)
				container.eventNext = time.Now().Add(timeToNextEvent)

			}
		}
	}
}

func (e *Theater) AddImageConfigurationUsedAsCache(container *Configuration) (err error) {
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
		id, err = container.Docker.ImageFindIdByName(container.Docker.GetImageName())
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

		var buildPath = container.Docker.GetBuildFolderPath()
		if buildPath != "" {
			err = container.Docker.ImageBuildFromFolder()
		} else {
			err = container.Docker.ImageBuildFromServer()
		}
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
	e.sceneCaos = append(e.sceneCaos, container)

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

		var buildPath = container.Docker.GetBuildFolderPath()
		if buildPath != "" {
			err = container.Docker.ImageBuildFromFolder()
		} else {
			err = container.Docker.ImageBuildFromServer()
		}
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

func (e *Theater) AddContainerAddContainerConfigurationToPrologueScene(container *Configuration) (err error) {
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

	if e.refCaos == nil {
		e.refCaos = make([]*Configuration, 0)
	}

	for _, container := range e.sceneCaos {

		if container.Docker.GetInitialized() == false {
			err = container.Docker.Init()
			if err != nil {
				util.TraceToLog()
				return
			}
		}

		container.containerStarted = true

		if container.Docker.GetContainerIsStarted() == true {
			continue
		}

		err = container.Docker.ContainerStartAfterBuild()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	for _, container := range e.sceneCaos {
		e.refCaos = append(e.refCaos, container)
	}

	return
}

func (e *Theater) buildAll() (err error) {
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
	if makeLabel == true {
		_, err = file.Write([]byte("reading time\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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

	// Linux specific stats, not populated on Windows.
	// Current is the number of pids in the cgroup
	if makeLabel == true {
		_, err = file.Write([]byte("Linux specific stats, not populated on Windows. Current is the number of pids in the cgroup\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Linux specific stats, not populated on Windows. Limit is the hard limit on the number of pids in the cgroup. A \"Limit\" of 0 means that there is no limit.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Total CPU time consumed. (Units: nanoseconds on Linux, Units: 100's of nanoseconds on Windows)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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

	// Total CPU time consumed per core (Linux). Not used on Windows.
	// Total CPU length.
	//if makeLabel == true {
	//	_, err = file.Write([]byte("Total length of CPUs used in time consumed per core fields (Linux). Not used on Windows.\t"))
	//	if err != nil {
	//		util.TraceToLog()
	//		return
	//	}
	//} else {
	//	_, err = file.Write([]byte(fmt.Sprintf("%v", len(stats.CPUStats.CPUUsage.PercpuUsage))))
	//	if err != nil {
	//		util.TraceToLog()
	//		return
	//	}
	//
	//	_, err = file.Write([]byte("\t"))
	//	if err != nil {
	//		util.TraceToLog()
	//		return
	//	}
	//}

	// Total CPU time consumed per core (Linux). Not used on Windows.
	// Units: nanoseconds.
	if len(stats.CPUStats.CPUUsage.PercpuUsage) != 0 {
		e.cpus = len(stats.CPUStats.CPUUsage.PercpuUsage)
	}

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

	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	if makeLabel == true {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux). Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows.Not populated for Hyper-V Containers.).\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux). Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers).\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("System Usage. Linux only.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Online CPUs. Linux only.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods with throttling active.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Throttling Data. Linux only. Number of periods when the container hits its throttling limit.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Throttling Data. Linux only. Aggregate time the container was throttled for in nanoseconds.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Total CPU time consumed. (Units: nanoseconds on Linux. Units: 100's of nanoseconds on Windows)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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

	// CPU Usage. Linux and Windows.
	// Total CPU time consumed per core (Linux). Not used on Windows.
	// Total length of CPUs
	//if makeLabel == true {
	//	_, err = file.Write([]byte("Total CPU time consumed per core (Units: nanoseconds on Linux). Not used on Windows. Length of CPUs\t"))
	//	if err != nil {
	//		util.TraceToLog()
	//		return
	//	}
	//} else {
	//	_, err = file.Write([]byte(fmt.Sprintf("%v", len(stats.PreCPUStats.CPUUsage.PercpuUsage))))
	//	if err != nil {
	//		util.TraceToLog()
	//		return
	//	}
	//
	//	_, err = file.Write([]byte("\t"))
	//	if err != nil {
	//		util.TraceToLog()
	//		return
	//	}
	//}

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

	for cpuNumber, cpuTime := range stats.PreCPUStats.CPUUsage.PercpuUsage {
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

	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
	if makeLabel == true {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in kernel mode (Units: nanoseconds on Linux) - Time spent by all container processes in kernel mode (Units: 100's of nanoseconds on Windows - Not populated for Hyper-V Containers.)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Time spent by tasks of the cgroup in user mode (Units: nanoseconds on Linux) - Time spent by all container processes in user mode (Units: 100's of nanoseconds on Windows. Not populated for Hyper-V Containers)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("System Usage. (Linux only)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Online CPUs. (Linux only)\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Aggregate time the container was throttled for in nanoseconds.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Number of periods with throttling active.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("Throttling Data. (Linux only) - Number of periods when the container hits its throttling limit.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("current res_counter usage for memory\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("maximum usage ever recorded.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("number of times memory usage hits limits.\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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

	if makeLabel == true {
		_, err = file.Write([]byte("memory limit\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("committed bytes\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("peak committed bytes\t"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
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
	if makeLabel == true {
		_, err = file.Write([]byte("private working set\n"))
		if err != nil {
			util.TraceToLog()
			return
		}
	} else {
		_, err = file.Write([]byte(fmt.Sprintf("%v", stats.MemoryStats.PrivateWorkingSet)))
		if err != nil {
			util.TraceToLog()
			return
		}

		_, err = file.Write([]byte("\n"))
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *Theater) getProbalityNumber() (probality float64) {
	return 1.0 - e.getRandSeed().Float64()
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
