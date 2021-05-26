package iotmakerdockerbuilder

// SetDockerfileBuilder
//
// English: Defines a new object containing the builder of the dockerfile.
//
//   Note: see the DockerfileAuto interface for further instructions.
//
// Português: Define um novo objeto contendo o construtor do arquivo dockerfile.
//
//   Nota: veja a interface DockerfileAuto para mais instruções.
func (e *ContainerBuilder) SetDockerfileBuilder(value DockerfileAuto) {
	e.autoDockerfile = value
}
