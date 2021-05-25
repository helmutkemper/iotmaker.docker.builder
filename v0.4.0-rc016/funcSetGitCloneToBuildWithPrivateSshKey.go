package iotmakerdockerbuilder

// SetGitCloneToBuildWithPrivateSSHKey (english):
//
// SetGitCloneToBuildWithPrivateSSHKey (português): Define o caminho de um repositório privado para ser usado como base
// da imagem a ser montada.
//
//   url: Endereço do repositório contendo o projeto
//   privateSSHKeyPath: este é o caminho da chave ssh privada compatível com a chave pública cadastrada no git
//   password: senha usada no momento que a chave ssh foi gerada
//
//     Nota: * Password é o password usado na criação da chave e não password do repositório;
//           o diretório ~/.ssh só é acessado pelo programa se o mesmo for inicializado com nível de acesso
//           administrador;
//           * Esta função deve ser usada com a função ImageBuildFromServer() e SetImageName() para baixar e gerar uma
//           imagem a partir do conteúdo de um repositório git;
//           * O repositório deve contar um arquivo Dockerfile e ele será procurado na seguinte ordem:
//           './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*', '.*Dockerfile.*' e
//           '.*dockerfile.*'
//
//
//     var err error
//     var usr *user.User
//     var privateSSHKeyPath string
//     var file []byte
//     usr, err = user.Current()
//     if err != nil {
//       panic(err)
//     }
//
//     privateSSHKeyPath = filepath.Join(usr.HomeDir, ".shh", "id_rsa")
//
//     var container = ContainerBuilder{}
//     container.SetGitCloneToBuildWithPrivateSSHKey(url, privateSSHKeyPath, password)
//     container.SetGitConfigFile(string(file))
func (e *ContainerBuilder) SetGitCloneToBuildWithPrivateSSHKey(url, privateSSHKeyPath, password string) {
	e.gitData.url = url
	e.gitData.sshPrivateKeyPath = privateSSHKeyPath
	e.gitData.password = password
}
