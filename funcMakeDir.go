package iotmakerdockerbuilder

import (
	"github.com/helmutkemper/util"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func MakeDir(path string) (err error) {
	var pathList []string

	path = filepath.Dir(path)

	path, err = filepath.Abs(path)
	if err != nil {
		util.TraceToLog()
		return
	}

	pathList = strings.Split(path, string(filepath.Separator))

	var subPathActual string
	for _, subPath := range pathList {
		if subPath == "" {
			continue
		}

		subPathActual += string(filepath.Separator) + subPath

		if DirCheckExists(subPathActual) == false {
			err = os.Mkdir(subPathActual, fs.ModePerm)
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	}

	return
}
