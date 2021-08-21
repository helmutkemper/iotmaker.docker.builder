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

// ContainerBuilder
//
// English: Docker manager
//
// Português: Gerenciador de containers e imagens docker
type ContainerBuilder struct {
	enableCache             bool
	network                 isolatedNetwork.ContainerBuilderNetworkInterface
	dockerSys               iotmakerdocker.DockerSystem
	changePointer           chan iotmakerdocker.ContainerPullStatusSendToChannel
	onContainerReady        *chan bool
	onContainerInspect      *chan bool
	imageInstallExtras      bool
	imageCacheName          string
	imageName               string
	imageID                 string
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

	buildOptions types.ImageBuildOptions
}
