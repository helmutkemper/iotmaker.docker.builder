package iotmakerdockerbuilder

import "time"

func (e ContainerBuilder) ImageMakeCacheWithDefaultName(projectPath string, expirationDate time.Duration) (err error) {
	return e.ImageMakeCache(projectPath, "cache:latest", expirationDate)
}
