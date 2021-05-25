package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
)

// addImageBuildOptionsGitCredentials (english): call from SetPrivateRepositoryAutoConfig()
//
// addImageBuildOptionsGitCredentials (português): chamada por SetPrivateRepositoryAutoConfig()
func (e *ContainerBuilder) addImageBuildOptionsGitCredentials() (buildOptions types.ImageBuildOptions) {

	if buildOptions.BuildArgs == nil {
		e.buildOptions.BuildArgs = make(map[string]*string)
	}

	if e.contentGitConfigFile != "" {
		e.buildOptions.BuildArgs["GITCONFIG_FILE"] = &e.contentGitConfigFile
	}

	if e.contentKnownHostsFile != "" {
		e.buildOptions.BuildArgs["KNOWN_HOSTS_FILE"] = &e.contentKnownHostsFile
	}

	if e.contentIdRsaFile != "" {
		e.buildOptions.BuildArgs["SSH_ID_RSA_FILE"] = &e.contentIdRsaFile
	}

	if e.gitPathPrivateRepository != "" {
		e.buildOptions.BuildArgs["GIT_PRIVATE_REPO"] = &e.gitPathPrivateRepository
	}

	return
}