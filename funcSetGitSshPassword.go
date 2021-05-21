package iotmakerdockerbuilder

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
