package iotmakerdockerbuilder

// SetImageBuildOptionsTarget (english): Build the specified stage as defined inside the Dockerfile. See the multi-stage
// build docs for details.
//
//   See https://docs.docker.com/develop/develop-images/multistage-build/
//
// SetImageBuildOptionsTarget (português): Constrói o container a partir do estágio definido no arquivo Dockerfile.
// Veja a documentação de múltiplos estágios para mais detalhes.
//
//   See https://docs.docker.com/develop/develop-images/multistage-build/
func (e *ContainerBuilder) SetImageBuildOptionsTarget(value string) {
	e.buildOptions.Target = value
}
