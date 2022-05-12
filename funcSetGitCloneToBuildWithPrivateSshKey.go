package iotmakerdockerbuilder

// SetGitCloneToBuildWithPrivateSSHKey
//
// English:
//
//  Defines the path of a repository to be used as the base of the image to be mounted.
//
//   Input:
//     url: Address of the repository containing the project
//     privateSSHKeyPath: this is the path of the private ssh key compatible with the public key
//       registered in git
//     password: password used when the ssh key was generated or empty string
//
// Note:
//
//   * If the repository is private and the host computer has access to the git server, use
//     SetPrivateRepositoryAutoConfig() to copy the git credentials contained in ~/.ssh and the
//     settings of ~/.gitconfig automatically;
//   * To be able to access private repositories from inside the container, build the image in two or
//     more steps and in the first step, copy the id_rsa and known_hosts files to the /root/.ssh
//     folder, and the ~/.gitconfig file to the /root folder;
//   * The MakeDefaultDockerfileForMe() function make a standard dockerfile with the procedures above;
//   * If the ~/.ssh/id_rsa key is password protected, use the SetGitSshPassword() function to set the
//     password;
//   * If you want to define the files manually, use SetGitConfigFile(), SetSshKnownHostsFile() and
//     SetSshIdRsaFile() to define the files manually;
//   * This function must be used with the ImageBuildFromServer() and SetImageName() function to
//     download and generate an image from the contents of a git repository;
//   * The repository must contain a Dockerfile file and it will be searched for in the following
//     order:
//     './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*',
//     '.*Dockerfile.*' and '.*dockerfile.*';
//   * The repository can be defined by the methods SetGitCloneToBuild(),
//     SetGitCloneToBuildWithPrivateSshKey(), SetGitCloneToBuildWithPrivateToken() and
//     SetGitCloneToBuildWithUserPassworh().
//
// code:
//
//       var err error
//       var usr *user.User
//       var privateSSHKeyPath string
//       var userGitConfigPath string
//       var file []byte
//       usr, err = user.Current()
//       if err != nil {
//         panic(err)
//       }
//
//       privateSSHKeyPath = filepath.Join(usr.HomeDir, ".shh", "id_ecdsa")
//       userGitConfigPath = filepath.Join(usr.HomeDir, ".gitconfig")
//       file, err = ioutil.ReadFile(userGitConfigPath)
//
//       var container = ContainerBuilder{}
//       container.SetGitCloneToBuildWithPrivateSSHKey(url, privateSSHKeyPath, password)
//       container.SetGitConfigFile(string(file))
//
// Português:
//
//  Define o caminho de um repositório para ser usado como base da imagem a ser montada.
//
//   Entrada:
//     url: Endereço do repositório contendo o projeto
//     privateSSHKeyPath: este é o caminho da chave ssh privada compatível com a chave pública
//       cadastrada no git
//     password: senha usada no momento que a chave ssh foi gerada ou string em branco
//
// Nota:
//
//   * Caso o repositório seja privado e o computador hospedeiro tenha acesso ao servidor git, use
//     SetPrivateRepositoryAutoConfig() para copiar as credências do git contidas em ~/.ssh e as
//     configurações de ~/.gitconfig de forma automática;
//   * Para conseguir acessar repositórios privados de dentro do container, construa a imagem em duas
//     ou mais etapas e na primeira etapa, copie os arquivos id_rsa e known_hosts para a pasta
//     /root/.ssh e o arquivo .gitconfig para a pasta /root/;
//   * A função MakeDefaultDockerfileForMe() monta um dockerfile padrão com os procedimentos acima;
//   * Caso a chave ~/.ssh/id_rsa seja protegida com senha, use a função SetGitSshPassword() para
//     definir a senha da mesma;
//   * Caso queira definir os arquivos de forma manual, use SetGitConfigFile(), SetSshKnownHostsFile()
//     e SetSshIdRsaFile() para definir os arquivos de forma manual;
//   * Esta função deve ser usada com a função ImageBuildFromServer() e SetImageName() para baixar e
//     gerar uma imagem a partir do conteúdo de um repositório git;
//   * O repositório deve contar um arquivo Dockerfile e ele será procurado na seguinte ordem:
//     './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*',
//     '.*Dockerfile.*' e '.*dockerfile.*';
//   * O repositório pode ser definido pelos métodos SetGitCloneToBuild(),
//     SetGitCloneToBuildWithPrivateSshKey(), SetGitCloneToBuildWithPrivateToken() e
//     SetGitCloneToBuildWithUserPassworh().
//
// code:
//
//       var err error
//       var usr *user.User
//       var privateSSHKeyPath string
//       var userGitConfigPath string
//       var file []byte
//       usr, err = user.Current()
//       if err != nil {
//         panic(err)
//       }
//
//       privateSSHKeyPath = filepath.Join(usr.HomeDir, ".shh", "id_ecdsa")
//       userGitConfigPath = filepath.Join(usr.HomeDir, ".gitconfig")
//       file, err = ioutil.ReadFile(userGitConfigPath)
//
//       var container = ContainerBuilder{}
//       container.SetGitCloneToBuildWithPrivateSSHKey(url, privateSSHKeyPath, password)
//       container.SetGitConfigFile(string(file))
func (e *ContainerBuilder) SetGitCloneToBuildWithPrivateSSHKey(url, privateSSHKeyPath, password string) {
	e.gitData.url = url
	e.gitData.sshPrivateKeyPath = privateSSHKeyPath
	e.gitData.password = password
}
