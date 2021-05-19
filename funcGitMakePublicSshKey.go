package iotmakerdockerbuilder

import (
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
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