package iotmakerdockerbuilder

import (
	"bytes"
	"github.com/helmutkemper/util"
)

func (e *ContainerBuilder) GetLastLineOfOccurrenceBySearchTextInsideContainerLog(value string) (text string, contains bool, err error) {
	var logs []byte
	var lineList [][]byte
	logs, err = e.GetContainerLog()
	if err != nil {
		util.TraceToLog()
		return
	}

	logs = bytes.ReplaceAll(logs, []byte("\r"), []byte(""))
	lineList = bytes.Split(logs, []byte("\n"))

	for i := len(lineList) - 1; i >= 0; i -= 1 {
		if bytes.Contains(lineList[i], []byte(value)) == true {
			text = string(lineList[i])
			contains = true
			return
		}
	}

	return
}
