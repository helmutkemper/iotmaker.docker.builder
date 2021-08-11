package iotmakerdockerbuilder

func (e *ContainerBuilder) MakeDefaultDockerfileForMeWithInstallExtras() {
	e.makeDefaultDockerfile = true
	e.imageInstallExtras = true
}
