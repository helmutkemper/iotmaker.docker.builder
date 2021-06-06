package iotmakerdockerbuilder

// SetImageBuildOptionsPlatform
//
// English: Target platform containers for this service will run on, using the os[/arch[/variant]] syntax, e.g.
//
//   osx
//   windows/amd64
//   linux/arm64/v8
//
// Português: Especifica a plataforma de container onde o serviço vai rodar, usando a sintaxe os[/arch[/variant]]
//
//   osx
//   windows/amd64
//   linux/arm64/v8
func (e *ContainerBuilder) SetImageBuildOptionsPlatform(value string) {
	e.buildOptions.Platform = value
}
