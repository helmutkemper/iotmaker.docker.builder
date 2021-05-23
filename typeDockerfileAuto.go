package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types/mount"
	dockerfileGolang "github.com/helmutkemper/iotmaker.docker.builder.golang.dockerfile"
)

// DockerfileAuto (english):
//
// DockerfileAuto (português):
type DockerfileAuto interface {
	MountDefaultDockerfile(args map[string]*string, changePorts []dockerfileGolang.ChangePort, openPorts []string, volumes []mount.Mount) (dockerfile string, err error)
}
