package iotmaker_docker_builder

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"io/ioutil"
	"os"
)

func (e *ContainerBuilder) gitMakePublicSshKey() (publicKeys *ssh.PublicKeys, err error) {
	_, err = os.Stat(e.gitData.sshPrivateKeyPath)
	if err != nil {
		return
	}

	publicKeys, err = ssh.NewPublicKeysFromFile("git", e.gitData.sshPrivateKeyPath, e.gitData.password)
	if err != nil {
		return
	}

	return
}

func (e *ContainerBuilder) ImageBuildFromServer() (err error) {
	err = e.verifyImageName()
	if err != nil {
		return
	}

	if e.gitData.user != "" && e.gitData.password == "" {
		err = errors.New("user is set, but, password is not set")
		return
	} else if e.gitData.user == "" && e.gitData.sshPrivateKeyPath == "" && e.gitData.password != "" {
		err = errors.New("password is set. now, set user or private ssh toke path")
		return
	}

	var tmpDirPath string
	var publicKeys *ssh.PublicKeys
	var gitCloneConfig *git.CloneOptions

	if e.gitData.sshPrivateKeyPath != "" {
		publicKeys, err = e.gitMakePublicSshKey()
		if err != nil {
			return
		}
	}

	tmpDirPath, err = ioutil.TempDir(os.TempDir(), "iotmaker.docker.builder.git.*")
	if err != nil {
		return
	}

	defer os.RemoveAll(tmpDirPath)

	if e.gitData.sshPrivateKeyPath != "" {
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

	e.imageID, err = e.dockerSys.ImageBuildFromFolder(
		tmpDirPath,
		[]string{
			e.imageName,
		},
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