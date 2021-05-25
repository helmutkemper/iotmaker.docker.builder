package iotmakerdockerbuilder

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

// SetPrivateRepositoryAutoConfig (english):
//
// SetPrivateRepositoryAutoConfig (português): Copia a chave ssh  ~/.ssh/id_rsa e o arquivo ~/.gitconfig para as
// variáveis SSH_ID_RSA_FILE e GITCONFIG_FILE.
//
// Dentro do Dockerfile
func (e *ContainerBuilder) SetPrivateRepositoryAutoConfig() (err error) {
	var userData *user.User
	var fileData []byte

	userData, err = user.Current()
	if err != nil {
		return
	}

	var filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "id_rsa")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}

	e.contentIdRsaFile = string(fileData)
	e.contentIdRsaFileWithScape = strings.ReplaceAll(e.contentIdRsaFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "known_hosts")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}

	e.contentKnownHostsFile = string(fileData)
	e.contentKnownHostsFileWithScape = strings.ReplaceAll(e.contentKnownHostsFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".gitconfig")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		return
	}

	e.contentGitConfigFile = string(fileData)
	e.contentGitConfigFileWithScape = strings.ReplaceAll(e.contentGitConfigFile, `"`, `\"`)

	e.addImageBuildOptionsGitCredentials()
	return
}
