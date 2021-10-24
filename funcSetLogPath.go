package iotmakerdockerbuilder

import "os"

func (e *ContainerBuilder) SetCsvLogPath(path string, removeOldFile bool) {

	if removeOldFile == true {
		_ = os.Remove(path)
	}

	e.chaos.logPath = path
}
