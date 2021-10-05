package iotmakerdockerbuilder

import (
	"bytes"
)

func (e *ContainerBuilder) logsCleaner(logs []byte) [][]byte {

	size := len(logs)

	// faz o log só lê a parte mais recente do mesmo
	logs = logs[e.logsLastSize:]
	e.logsLastSize = size

	logs = bytes.ReplaceAll(logs, []byte("\r"), []byte(""))
	return bytes.Split(logs, []byte("\n"))
}
