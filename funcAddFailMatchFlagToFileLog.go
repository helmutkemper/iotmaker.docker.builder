package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"path/filepath"
	"strings"
)

// AddFailMatchFlagToFileLog
//
// Similar: AddFailMatchFlag(), AddFailMatchFlagToFileLog(), AddFilterToFail()
//
// English: Error text searched for in the container's standard output.
//   Input:
//     value: Error text
//     logFilePath: path to diretory to save container default output into file
//   Output:
//     err: Default error object
//
// Português: Texto indicativo de erro procurado na saída padrão do container.
//   Entrada:
//     value: Texto indicativo de erro
//     logFilePath: caminho do diretório para salvar a saída padrão do container em arquivo
//   Output:
//     err: Objeto de erro padrão
func (e *ContainerBuilder) AddFailMatchFlagToFileLog(value, logDirectoryPath string) (err error) {
	if e.chaos.filterFail == nil {
		e.chaos.filterFail = make([]LogFilter, 0)
	}

	if strings.HasPrefix(logDirectoryPath, string(filepath.Separator)) == false {
		logDirectoryPath += string(filepath.Separator)
	}

	err = util.DirMake(logDirectoryPath)
	if err != nil {
		util.TraceToLog()
		return
	}

	e.chaos.filterFail = append(e.chaos.filterFail, LogFilter{Match: value, LogPath: logDirectoryPath})

	return
}
