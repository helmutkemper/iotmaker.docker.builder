package iotmakerdockerbuilder

// MakeDefaultDockerfileForMe (english):
//
// MakeDefaultDockerfileForMe (português): Monta o arquivo Dockerfile-iotmaker dentro da pasta alvo de forma
// automática.
//
// Caso haja portas expostas nas configurações, as mesmas serão definidas automaticamente e o mesmo vale para volumes,
// onde arquivos compartilhados entre o host e o container irá expor o volume contendo o arquivo dentro do container.
//
//   Cuidado: o arquivo Dockerfile-iotmaker pode ser sobrescrito.
//
//   Regras: Para projetos Golang, o arquivo go.mod é obrigatório
//   o arquivo main.go contendo o package main deve está na raiz do diretório
//
func (e *ContainerBuilder) MakeDefaultDockerfileForMe() {
	e.makeDefaultDockerfile = true
}
