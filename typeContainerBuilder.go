package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	dockerfileGolang "github.com/helmutkemper/iotmaker.docker.builder.golang.dockerfile"
	isolatedNetwork "github.com/helmutkemper/iotmaker.docker.builder.network.interface"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"time"
)

type chaos struct {
	filterToStart              []LogFilter
	filterRestart              []LogFilter
	filterSuccess              []LogFilter
	filterFail                 []LogFilter
	filterLog                  []LogFilter
	sceneName                  string
	logPath                    string
	serviceStartedAt           time.Time
	minimumTimeBeforeRestart   time.Duration
	maximumTimeBeforeRestart   time.Duration
	minimumTimeToStartChaos    time.Duration
	maximumTimeToStartChaos    time.Duration
	minimumTimeToPause         time.Duration
	maximumTimeToPause         time.Duration
	minimumTimeToUnpause       time.Duration
	maximumTimeToUnpause       time.Duration
	minimumTimeToRestart       time.Duration
	maximumTimeToRestart       time.Duration
	restartProbability         float64
	restartChangeIpProbability float64
	restartLimit               int
	enableChaos                bool
	event                      chan Event
	monitorStop                chan struct{}
	monitorRunning             bool
	//containerStarted         bool
	containerPaused          bool
	containerStopped         bool
	linear                   bool
	chaosStarted             bool
	chaosCanRestartContainer bool
	//chaosCanRestartEnd       bool
	eventNext time.Time
}

// ContainerBuilder
//
// English: Docker manager
//
// Português: Gerenciador de containers e imagens docker
type ContainerBuilder struct {
	csvValueSeparator       string
	csvRowSeparator         string
	csvConstHeader          bool
	logCpus                 int
	rowsToPrint             int64
	chaos                   chaos
	enableCache             bool
	network                 isolatedNetwork.ContainerBuilderNetworkInterface
	dockerSys               iotmakerdocker.DockerSystem
	changePointer           chan iotmakerdocker.ContainerPullStatusSendToChannel
	onContainerReady        *chan bool
	onContainerInspect      *chan bool
	imageInspected          bool
	imageInstallExtras      bool
	imageCacheName          string
	imageName               string
	imageID                 string
	imageRepoTags           []string
	imageRepoDigests        []string
	imageParent             string
	imageComment            string
	imageCreated            time.Time
	imageContainer          string
	imageAuthor             string
	imageArchitecture       string
	imageVariant            string
	imageOs                 string
	imageOsVersion          string
	imageSize               int64
	imageVirtualSize        int64
	containerName           string
	buildPath               string
	environmentVar          []string
	changePorts             []dockerfileGolang.ChangePort
	openPorts               []string
	exposePortsOnDockerfile []string
	openAllPorts            bool
	waitString              string
	waitStringTimeout       time.Duration
	containerID             string
	ticker                  *time.Ticker
	inspect                 iotmakerdocker.ContainerInspect
	logs                    string
	logsLastSize            int
	inspectInterval         time.Duration
	gitData                 gitData
	volumes                 []mount.Mount
	IPV4Address             string
	autoDockerfile          DockerfileAuto
	containerConfig         container.Config
	restartPolicy           iotmakerdocker.RestartPolicy

	makeDefaultDockerfile bool
	printBuildOutput      bool
	init                  bool
	startedAfterBuild     bool

	contentIdRsaFile               string
	contentIdRsaFileWithScape      string
	contentKnownHostsFile          string
	contentKnownHostsFileWithScape string
	contentGitConfigFile           string
	contentGitConfigFileWithScape  string

	gitPathPrivateRepository string

	buildOptions        types.ImageBuildOptions
	imageExpirationTime time.Duration
}
