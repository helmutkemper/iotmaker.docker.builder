package iotmakerdockerbuilder

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"

// MakeDefaultDockerfileForMe
//
// English: Automatically mount the Dockerfile-iotmaker inside the target folder.
//
// If there are ports exposed in the configurations, they will be defined automatically and the same goes for
// volumes, where files shared between the host and the container will expose the folder containing the file inside
// the container as volume.
//
//   Caution: the Dockerfile-iotmaker may be overwritten.
//
//   Rules: For Golang projects, the go.mod file is mandatory;
//   The main.go file containing the main package must be at the root folder.
//
//   Note: If you need a dockerfile made for another programming language, see the DockerfileAuto interface and the
//   SetDockerfileBuilder() function
//
// Português: Monta o arquivo Dockerfile-iotmaker dentro da pasta alvo de forma automática.
//
// Caso haja portas expostas nas configurações, as mesmas serão definidas automaticamente e o mesmo vale para
// volumes, onde arquivos compartilhados entre o host e o container exporá a pasta contendo o arquivo dentro do
// container como um volume.
//
//   Cuidado: o arquivo Dockerfile-iotmaker pode ser sobrescrito.
//
//   Regras: Para projetos Golang, o arquivo go.mod é obrigatório;
//   O arquivo main.go contendo o package main deve está na raiz do diretório.
//
//   Nota: Caso necessite de um dockerfile feito para outra linguagem de programação, veja a interface
//   DockerfileAuto e a função SetDockerfileBuilder()
func (e *ContainerBuilder) MakeDefaultDockerfileForMe() {

	if e.enableCache == true {

		// se a cache foi habilitada e não existe imagem cache, crie o dockerfile completo
		// isto foi feito devido a falhas nos testes com usuários
		var id string
		var dockerSys = iotmakerdocker.DockerSystem{}
		_ = dockerSys.Init()
		id, _ = dockerSys.ImageFindIdByName(e.imageCacheName)
		if id == "" {
			e.imageInstallExtras = true
		}
	}

	e.makeDefaultDockerfile = true
}
