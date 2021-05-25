package iotmakerdockerbuilder

import (
	"bytes"
)

// FindTextInsideContainerLog
//
// English: search for text in standard container output.
//
//   value: searched text
//
// Português: procurar por um texto na saída padrão do container.
//
//   value: texto procurado
func (e *ContainerBuilder) FindTextInsideContainerLog(value string) (contains bool, err error) {
	var logs []byte
	logs, err = e.GetContainerLog()
	if err != nil {
		return
	}

	contains = bytes.Contains(logs, []byte(value))
	return
}
