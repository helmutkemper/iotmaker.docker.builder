package iotmakerdockerbuilder

import (
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"os"
)

// gitMakePublicSshKey (english):
//
// gitMakePublicSshKey (português):
func (e *ContainerBuilder) gitMakePublicSshKey() (publicKeys *ssh.PublicKeys, err error) {
	if e.gitData.sshPrivateKeyPath != "" {
		_, err = os.Stat(e.gitData.sshPrivateKeyPath)
		if err != nil {
			return
		}
		publicKeys, err = ssh.NewPublicKeysFromFile("git", e.gitData.sshPrivateKeyPath, e.gitData.password)
	} else if e.contentIdRsaFile != "" {
		publicKeys, err = ssh.NewPublicKeys("git", []byte(e.contentIdRsaFile), e.gitData.password)
	}

	if err != nil {
		return
	}

	return
}
