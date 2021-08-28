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
	"strconv"
	"time"
)

type Theater struct {
	sceneCache    []*Configuration
	sceneBuilding []*Configuration
	scenePrologue []*Configuration
	sceneLinear   []*Configuration
	sceneCaos     []*Configuration

	refLinear []*Configuration
	refCaos   []*Configuration

	ticker *time.Ticker
}

type Timers struct {
	Min time.Duration
	Max time.Duration
}

type LogFilter struct {
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
	FilterToStart      []LogFilter //todo: se não tiver filtro, restart pot tempo aleatório
	TimeToStart        Timers
	RestartProbability float64
	RestartLimit       int
}

type Caos struct {
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

	Caos Caos

	Linear            bool
	caosStarted       bool
	caosCanRestart    bool
	caosCanRestartEnd bool

	containerStarted bool
	containerPaused  bool
	containerStopped bool

	testEnd bool

	err chan error

	eventNext time.Time
}

func NewContainer(docker *dockerBuild.ContainerBuilder) (configuration *Configuration) {
	return &Configuration{Docker: docker}
}

func (e *Configuration) SetLinear() (configuration *Configuration) {
	e.Linear = true
	return e
}

func (e *Configuration) SetLogPath(path string) (configuration *Configuration) {
	e.LogPath = path
	return e
}

func (e *Configuration) AddFilterToLog(match, filter, search, replace string) (configuration *Configuration) {

	if e.Log == nil {
		e.Log = make([]LogFilter, 0)
	}

	e.Log = append(
		e.Log, LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

func (e *Configuration) AddFilterToFailEvent(match, filter, search, replace string) (configuration *Configuration) {

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

func (e *Configuration) AddFilterToEndEvent(match, filter, search, replace string) (configuration *Configuration) {

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

func (e *Configuration) AddFilterToStartCaos(match, filter, search, replace string) (configuration *Configuration) {

	if e.Caos.FilterToStart == nil {
		e.Caos.FilterToStart = make([]LogFilter, 0)
	}

	e.Caos.FilterToStart = append(
		e.Caos.FilterToStart, LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

func (e *Configuration) AddCaosPauseDuration(min, max time.Duration) (configuration *Configuration) {

	e.Caos.TimeToPause.Min = min
	e.Caos.TimeToPause.Max = max

	if e.Caos.TimeToStart.Min == 0 && e.Caos.TimeToStart.Max == 0 {
		e.Caos.TimeToStart.Min = min
		e.Caos.TimeToStart.Max = max
	}

	if e.Caos.TimeToUnpause.Min == 0 && e.Caos.TimeToUnpause.Max == 0 {
		e.Caos.TimeToUnpause.Min = min
		e.Caos.TimeToUnpause.Max = max
	}

	return e
}

func (e *Configuration) AddCaosUnpauseDuration(min, max time.Duration) (configuration *Configuration) {

	e.Caos.TimeToUnpause.Min = min
	e.Caos.TimeToUnpause.Max = max

	if e.Caos.TimeToStart.Min == 0 && e.Caos.TimeToStart.Max == 0 {
		e.Caos.TimeToStart.Min = min
		e.Caos.TimeToStart.Max = max
	}

	if e.Caos.TimeToPause.Min == 0 && e.Caos.TimeToPause.Max == 0 {
		e.Caos.TimeToPause.Min = min
		e.Caos.TimeToPause.Max = max
	}

	return e
}

func (e *Configuration) AddCaosStartDuration(min, max time.Duration) (configuration *Configuration) {

	e.Caos.TimeToStart.Min = min
	e.Caos.TimeToStart.Max = max

	if e.Caos.TimeToUnpause.Min == 0 && e.Caos.TimeToUnpause.Max == 0 {
		e.Caos.TimeToUnpause.Min = min
		e.Caos.TimeToUnpause.Max = max
	}

	if e.Caos.TimeToPause.Min == 0 && e.Caos.TimeToPause.Max == 0 {
		e.Caos.TimeToPause.Min = min
		e.Caos.TimeToPause.Max = max
	}

	return e
}

func (e *Configuration) AddCaosRestartFilter(match, filter, search, replace string) (configuration *Configuration) {

	if e.Caos.Restart == nil {
		e.Caos.Restart = &Restart{}
	}

	if e.Caos.Restart.FilterToStart == nil {
		e.Caos.Restart.FilterToStart = make([]LogFilter, 0)
	}

	e.Caos.Restart.FilterToStart = append(
		e.Caos.Restart.FilterToStart, LogFilter{
			Match:   match,
			Filter:  filter,
			Search:  search,
			Replace: replace,
		},
	)

	return e
}

func (e *Configuration) AddCaosRestartInterval(min, max time.Duration) (configuration *Configuration) {

	if e.Caos.Restart == nil {
		e.Caos.Restart = &Restart{}
	}

	e.Caos.Restart.TimeToStart.Min = min
	e.Caos.Restart.TimeToStart.Max = max

	return e
}

func (e *Configuration) AddCaosRestartController(probability float64, limit int) (configuration *Configuration) {

	if e.Caos.Restart == nil {
		e.Caos.Restart = &Restart{}
	}

	e.Caos.Restart.RestartProbability = probability
	e.Caos.Restart.RestartLimit = limit

	return e
}

func (e *Theater) Init() {
	e.ticker = time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-e.ticker.C:
				e.manager()
			}
		}
	}()
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

	var inspect iotmakerdocker.ContainerInspect

	for _, container := range e.refCaos {

		probality = e.getProbalityNumber()

		inspect, err = container.Docker.ContainerInspect()
		if err != nil {
			util.TraceToLog()
			container.err <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
			continue
		}

		if inspect.State.OOMKilled == true {
			container.err <- errors.New(container.Docker.GetContainerName() + ".error: OOMKilled")
			continue
		}

		if inspect.State.Dead == true {
			container.err <- errors.New(container.Docker.GetContainerName() + ".error: dead")
			continue
		}

		if inspect.State.ExitCode != 0 {
			container.err <- errors.New(container.Docker.GetContainerName() + ".error: exit code " + strconv.Itoa(inspect.State.ExitCode))
			continue
		}

		if (container.containerStopped == true || container.containerPaused == true) != true {
			if inspect.State.Running == false {
				container.err <- errors.New(container.Docker.GetContainerName() + ".error: not running")
				continue
			}

			if inspect.State.Paused == false {
				container.err <- errors.New(container.Docker.GetContainerName() + ".error: paused")
				continue
			}

			if inspect.State.Restarting == false {
				container.err <- errors.New(container.Docker.GetContainerName() + ".error: restarting")
				continue
			}
		}

		logs, err = container.Docker.GetContainerLog()
		if err != nil {
			container.err <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
			util.TraceToLog()
			continue
		}

		lineList = e.logsCleaner(logs)

		err = e.writeContainerLogToFile(container.LogPath, lineList, container)
		if err != nil {
			container.err <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
			util.TraceToLog()
			continue
		}

		if container.Linear == true {
			continue
		}

		// flag indicando que o container pode ser reiniciado
		if container.Caos.Restart != nil {
			_, found = e.logsSearchSimplesText(lineList, container.Caos.Restart.FilterToStart)
			if container.caosCanRestart == false {
				if found == true {
					container.caosCanRestart = true
				}
			}
		}

		// flag o caos pode ser inicializado
		_, found = e.logsSearchSimplesText(lineList, container.Caos.FilterToStart)
		if container.caosStarted == false && found == true {
			container.caosStarted = true
			timeToNextEvent = e.selectBetweenMaxAndMin(container.Caos.TimeToStart.Max, container.Caos.TimeToStart.Min)
			container.eventNext = time.Now().Add(timeToNextEvent)
		}

		if container.containerStarted == false {
			continue
		}

		line, found = e.logsSearchSimplesText(lineList, container.Fail)
		if found == true {
			container.err <- errors.New(container.Docker.GetContainerName() + ".error: test fail - " + string(line))
		}

		line, found = e.logsSearchSimplesText(lineList, container.End)
		if found == true {
			container.testEnd = true
		}

		if time.Now().After(container.eventNext) == true || time.Now().Equal(container.eventNext) == true {

			if container.containerPaused == true {

				err = container.Docker.ContainerUnpause()
				if err != nil {
					container.err <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					util.TraceToLog()
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Caos.TimeToPause.Max, container.Caos.TimeToPause.Min)

			} else if container.containerStopped == true {

				err = container.Docker.ContainerStart()
				if err != nil {
					container.err <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					util.TraceToLog()
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Caos.TimeToPause.Max, container.Caos.TimeToPause.Min)

			} else if container.caosCanRestart == true && container.Caos.Restart != nil && container.Caos.Restart.RestartProbability <= probality && container.Caos.Restart.RestartLimit > 0 {

				err = container.Docker.ContainerStop()
				if err != nil {
					container.err <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					util.TraceToLog()
					continue
				}
				container.Caos.Restart.RestartLimit -= 1
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Caos.TimeToStart.Max, container.Caos.TimeToStart.Min)

			} else {

				err = container.Docker.ContainerPause()
				if err != nil {
					container.err <- errors.New(container.Docker.GetContainerName() + ".error: " + err.Error())
					util.TraceToLog()
					continue
				}
				timeToNextEvent = e.selectBetweenMaxAndMin(container.Caos.TimeToUnpause.Max, container.Caos.TimeToUnpause.Min)

			}

			container.eventNext = time.Now().Add(timeToNextEvent)
		}
	}
}

func (e *Theater) AddCache(container *Configuration) (err error) {
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

func (e *Theater) AddContainers(container *Configuration) (err error) {
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

func (e *Theater) buildContainers() (err error) {
	var id string

	for _, container := range e.sceneBuilding {
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

		err = container.Docker.ContainerBuildWithoutStartingItFromImage()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *Theater) AddContainerToPrologueScene(container *Configuration) (err error) {
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

func (e *Theater) AddContainerToLinearScene(container *Configuration) (err error) {
	if e.sceneLinear == nil {
		e.sceneLinear = make([]*Configuration, 0)
	}

	if container.Docker.GetInitialized() == true {
		err = errors.New("do not initialize the container")
		return
	}

	e.sceneLinear = append(e.sceneLinear, container)

	return
}

func (e *Theater) startLinearScene() (err error) {

	if e.refLinear == nil {
		e.refLinear = make([]*Configuration, 0)
	}

	for _, container := range e.sceneLinear {
		if container.Docker.GetContainerIsStarted() == true {
			continue
		}

		err = container.Docker.ContainerStartAfterBuild()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	for _, container := range e.sceneLinear {
		if container.LogPath == "" && container.Log == nil && container.Fail == nil && container.End == nil {
			continue
		}

		e.refLinear = append(e.refLinear, container)
	}

	return
}

func (e *Theater) AddContainerToCaosScene(container *Configuration) (err error) {
	if e.sceneCaos == nil {
		e.sceneCaos = make([]*Configuration, 0)
	}

	if container.Docker.GetInitialized() == true {
		err = errors.New("do not initialize the container")
		return
	}

	e.sceneCaos = append(e.sceneCaos, container)

	return
}

func (e *Theater) startCaosScene() (err error) {

	if e.refCaos == nil {
		e.refCaos = make([]*Configuration, 0)
	}

	for _, container := range e.sceneCaos {
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
		if container.LogPath == "" && container.Log == nil &&
			container.Fail == nil && container.End == nil {
			continue
		}

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

	err = e.startPrologueScene()
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

	var file *os.File
	file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		util.TraceToLog()
		return
	}

	defer file.Close()

	var skipMatch = make([]bool, len(configuration.Log))

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
					toFile = bytes.ReplaceAll(toFile, []byte(configuration.Log[filterLine].Search), []byte(configuration.Log[filterLine].Replace))
				}

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

	var stats = types.Stats{}
	stats, err = configuration.Docker.ContainerStatisticsOneShot()
	if err != nil {
		util.TraceToLog()
		return
	}

	// time
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

	// Linux specific stats, not populated on Windows.
	// Current is the number of pids in the cgroup
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

	// Linux specific stats, not populated on Windows.
	// Limit is the hard limit on the number of pids in the cgroup.
	// A "Limit" of 0 means that there is no limit.
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

	// Total CPU time consumed.
	// Units: nanoseconds (Linux)
	// Units: 100's of nanoseconds (Windows)
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

	// Total CPU time consumed per core (Linux). Not used on Windows.
	// Units: nanoseconds.
	_, err = file.Write([]byte(fmt.Sprintf("%v", stats.CPUStats.CPUUsage.PercpuUsage)))
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = file.Write([]byte("\t"))
	if err != nil {
		util.TraceToLog()
		return
	}

	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
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

	// Time spent by tasks of the cgroup in user mode (Linux).
	// Time spent by all container processes in user mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers
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

	// System Usage. Linux only.
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

	// Online CPUs. Linux only.
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

	// Throttling Data. Linux only.
	// Number of periods with throttling active
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

	// Throttling Data. Linux only.
	// Number of periods when the container hits its throttling limit.
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

	// Throttling Data. Linux only.
	// Aggregate time the container was throttled for in nanoseconds.
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

	// CPU Usage. Linux and Windows.
	// Total CPU time consumed.
	// Units: nanoseconds (Linux)
	// Units: 100's of nanoseconds (Windows)
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

	// CPU Usage. Linux and Windows.
	// Total CPU time consumed per core (Linux). Not used on Windows.
	// Units: nanoseconds.
	_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.CPUUsage.PercpuUsage)))
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = file.Write([]byte("\t"))
	if err != nil {
		util.TraceToLog()
		return
	}

	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Time spent by all container processes in kernel mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers.
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

	// CPU Usage. Linux and Windows.
	// Time spent by tasks of the cgroup in user mode (Linux).
	// Time spent by all container processes in user mode (Windows).
	// Units: nanoseconds (Linux).
	// Units: 100's of nanoseconds (Windows). Not populated for Hyper-V Containers
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

	// System Usage. Linux only.
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

	// Online CPUs. Linux only.
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

	// Throttling Data. Linux only.
	_, err = file.Write([]byte(fmt.Sprintf("%v", stats.PreCPUStats.ThrottlingData)))
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = file.Write([]byte("\t"))
	if err != nil {
		util.TraceToLog()
		return
	}

	// current res_counter usage for memory
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

	// maximum usage ever recorded.
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

	// number of times memory usage hits limits.
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

	// committed bytes
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

	// peak committed bytes
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

	// private working set
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
