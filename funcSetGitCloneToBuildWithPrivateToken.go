package iotmakerdockerbuilder

// SetGitCloneToBuildWithPrivateToken (english):
//
// SetGitCloneToBuildWithPrivateToken (português): Define o caminho de um repositório privado para ser usado como base
// da imagem a ser montada.
//
//   url: Endereço do repositório contendo o projeto
//   privateToken: Token gerado dentro do portal git
//
//     Nota: * Esta função deve ser usada com a função ImageBuildFromServer() e SetImageName() para baixar e gerar uma
//           imagem a partir do conteúdo de um repositório git;
//           * O repositório deve contar um arquivo Dockerfile e ele será procurado na seguinte ordem:
//           './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*', '.*Dockerfile.*' e
//           '.*dockerfile.*'
//
func (e *ContainerBuilder) SetGitCloneToBuildWithPrivateToken(url, privateToken string) {
	e.gitData.url = url
	e.gitData.privateToke = privateToken
}
