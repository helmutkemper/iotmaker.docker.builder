package iotmakerdockerbuilder

// SetGitCloneToBuildWithUserPassworh (english):
//
// SetGitCloneToBuildWithUserPassworh (português): Define o caminho de um repositório privado para ser usado como base
// da imagem a ser montada.
//
//   url: Endereço do repositório contendo o projeto
//   user: Usuário de acesso ao git
//   password: Senha de acesso ao git
//
//     Nota: * Esta função deve ser usada com a função ImageBuildFromServer() e SetImageName() para baixar e gerar uma
//           imagem a partir do conteúdo de um repositório git;
//           * O repositório deve contar um arquivo Dockerfile e ele será procurado na seguinte ordem:
//           './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*', '.*Dockerfile.*' e
//           '.*dockerfile.*'
//
func (e *ContainerBuilder) SetGitCloneToBuildWithUserPassworh(url, user, password string) {
	e.gitData.url = url
	e.gitData.user = user
	e.gitData.password = password
}

// SetGitSshPassword (english):
//
// SetGitSshPassword (português): Define a senha da chave ssh para repositórios git privados.
//
//   Nota: O repositório pode ser definido pelos métodos SetGitCloneToBuild(), SetGitCloneToBuildWithPrivateSshKey(),
//   SetGitCloneToBuildWithPrivateToken() e SetGitCloneToBuildWithUserPassworh()
//
//   SetPrivateRepositoryAutoConfig() copia as credencias do git contidas em ~/.ssh/id_rsa e as configurações de
//   ~/.gitconfig
//
//   Caso o certificado ssh seja protegido com chave, ela pode ser definida com SetGitSshPassword()
//
func (e *ContainerBuilder) SetGitSshPassword(password string) {
	e.gitData.password = password
}
