package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"path/filepath"
	"strings"
)

func (e *ContainerBuilder) AddRestartMatchFlagToFileLog(value, logDirectoryPath string) (err error) {
	if e.chaos.filterRestart == nil {
		e.chaos.filterRestart = make([]LogFilter, 0)
	}

	if strings.HasPrefix(logDirectoryPath, string(filepath.Separator)) == false {
		logDirectoryPath += string(filepath.Separator)
	}

	err = util.DirMake(logDirectoryPath)
	if err != nil {
		util.TraceToLog()
		return
	}

	e.chaos.filterRestart = append(e.chaos.filterRestart, LogFilter{Match: value, LogPath: logDirectoryPath})

	return
}
