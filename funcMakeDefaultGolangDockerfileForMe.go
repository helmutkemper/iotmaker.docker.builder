package iotmakerdockerbuilder

// MakeDefaultGolangDockerfileForMe (english):
//
// MakeDefaultGolangDockerfileForMe (português): Monta o arquivo Dockerfile-iotmaker dentro da pasta alvo de forma
// automática.
//
// Caso haja portas expostas nas configurações, as mesmas definidas automaticamente e o mesmo vale para volumes, onde
// arquivos compartilhados entre o host e o container irá expor o volume contendo o arquivo dentro do container.
//
//   Cuidado: o arquivo Dockerfile-iotmaker pode ser sobrescrito.
//
func (e *ContainerBuilder) MakeDefaultGolangDockerfileForMe() {
	e.makeDefaultDockerfile = true
}
