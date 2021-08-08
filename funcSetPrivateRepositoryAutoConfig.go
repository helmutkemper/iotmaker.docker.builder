package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

// SetPrivateRepositoryAutoConfig
//
// English: Copies the ssh ~/.ssh/id_rsa file and the ~/.gitconfig file to the SSH_ID_RSA_FILE and GITCONFIG_FILE
// variables.
//
// Português: Copia o arquivo ssh ~/.ssh/id_rsa e o arquivo ~/.gitconfig para as variáveis SSH_ID_RSA_FILE e
// GITCONFIG_FILE.
func (e *ContainerBuilder) SetPrivateRepositoryAutoConfig() (err error) {
	var userData *user.User
	var fileData []byte

	userData, err = user.Current()
	if err != nil {
		util.TraceToLog()
		return
	}

	var filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "id_rsa")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		util.TraceToLog()
		return
	}

	e.contentIdRsaFile = string(fileData)
	e.contentIdRsaFileWithScape = strings.ReplaceAll(e.contentIdRsaFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".ssh", "known_hosts")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		util.TraceToLog()
		return
	}

	e.contentKnownHostsFile = string(fileData)
	e.contentKnownHostsFileWithScape = strings.ReplaceAll(e.contentKnownHostsFile, `"`, `\"`)

	filePathToRead = filepath.Join(userData.HomeDir, ".gitconfig")
	fileData, err = ioutil.ReadFile(filePathToRead)
	if err != nil {
		util.TraceToLog()
		return
	}

	e.contentGitConfigFile = string(fileData)
	e.contentGitConfigFileWithScape = strings.ReplaceAll(e.contentGitConfigFile, `"`, `\"`)

	e.addImageBuildOptionsGitCredentials()
	return
}
