package iotmakerdockerbuilder

// SetGitPathPrivateRepository
//
// English: path do private repository defined in "go env -w GOPRIVATE=$GIT_PRIVATE_REPO"
//
//   Example: github.com/helmutkemper
//
// Português: caminho do repositório privado definido em "go env -w GOPRIVATE=$GIT_PRIVATE_REPO"
//
//   Exemplo: github.com/helmutkemper
//
func (e *ContainerBuilder) SetGitPathPrivateRepository(value string) {
	e.gitPathPrivateRepository = value
}
