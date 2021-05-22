package iotmakerdockerbuilder

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ImageBuildFromServer (english):
//
// ImageBuildFromServer (português): Monta uma imagem docker a partir de um projeto contido em um repositório git.
//
//   Nota: O repositório pode ser definido pelos métodos SetGitCloneToBuild(), SetGitCloneToBuildWithPrivateSshKey(),
//   SetGitCloneToBuildWithPrivateToken() e SetGitCloneToBuildWithUserPassworh()
//
//   SetPrivateRepositoryAutoConfig() copia as credências do git contidas em ~/.ssh/id_rsa e as configurações de
//   ~/.gitconfig
//
func (e *ContainerBuilder) ImageBuildFromServer() (err error) {
	err = e.verifyImageName()
	if err != nil {
		return
	}

	var tmpDirPath string
	var publicKeys *ssh.PublicKeys
	var gitCloneConfig *git.CloneOptions

	publicKeys, err = e.gitMakePublicSshKey()
	if err != nil {
		return
	}

	tmpDirPath, err = ioutil.TempDir(os.TempDir(), "iotmaker.docker.builder.git.*")
	if err != nil {
		return
	}

	defer os.RemoveAll(tmpDirPath)
	if e.gitData.sshPrivateKeyPath != "" || e.contentIdRsaFile != "" {
		gitCloneConfig = &git.CloneOptions{
			URL:      e.gitData.url,
			Auth:     publicKeys,
			Progress: nil,
		}
	} else if e.gitData.privateToke != "" {
		gitCloneConfig = &git.CloneOptions{
			// The intended use of a GitHub personal access token is in replace of your password
			// because access tokens can easily be revoked.
			// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
			Auth: &http.BasicAuth{
				Username: "abc123", // yes, this can be anything except an empty string
				Password: e.gitData.privateToke,
			},
			URL:      e.gitData.url,
			Progress: nil,
		}
	} else if e.gitData.user != "" && e.gitData.password != "" {
		gitCloneConfig = &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: e.gitData.user,
				Password: e.gitData.password,
			},
			URL:      e.gitData.url,
			Progress: nil,
		}
	} else {
		gitCloneConfig = &git.CloneOptions{
			URL:      e.gitData.url,
			Progress: nil,
		}
	}

	_, err = git.PlainClone(tmpDirPath, false, gitCloneConfig)
	if err != nil {
		return
	}

	if e.makeDefaultDockerfile == true {
		var dockerfile string
		var fileList []fs.FileInfo

		fileList, err = ioutil.ReadDir(tmpDirPath)
		log.Printf("6 %v", err)
		if err != nil {
			return
		}

		var pass = false
		for _, file := range fileList {
			if file.Name() == "go.mod" {
				pass = true
				break
			}
		}
		if pass == false {
			log.Printf("7 %v", err)
			err = errors.New("go.mod file not found")
			return
		}

		dockerfile, err = e.mountDefaultDockerfile()
		log.Printf("8 %v", err)
		if err != nil {
			return
		}

		var dockerfilePath = filepath.Join(tmpDirPath, "Dockerfile-iotmaker")
		log.Printf("9 %v", err)
		err = ioutil.WriteFile(dockerfilePath, []byte(dockerfile), os.ModePerm)
		if err != nil {
			return
		}
	}

	e.imageID, err = e.dockerSys.ImageBuildFromFolder(
		tmpDirPath,
		e.imageName,
		[]string{},
		e.buildOptions,
		e.changePointer,
	)

	if err != nil {
		return
	}

	if e.imageID == "" {
		err = errors.New("image ID was not generated")
		return
	}

	// Construir uma imagem de múltiplas etapas deixa imagens grandes e sem serventia, ocupando espaço no HD.
	err = e.dockerSys.ImageGarbageCollector()
	if err != nil {
		return
	}

	return
}
