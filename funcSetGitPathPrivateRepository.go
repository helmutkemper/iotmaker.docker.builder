package iotmakerdockerbuilder

// SetGitPathPrivateRepository (english):
//
//   Example: github.com/helmutkemper
//
// SetGitPathPrivateRepository (português):
//
//   Exemplo: github.com/helmutkemper
//
func (e *ContainerBuilder) SetGitPathPrivateRepository(value string) {
	e.gitPathPrivateRepository = value
}
