package iotmakerdockerbuilder

import "bytes"

func (e ContainerBuilder) logsCleaner(logs []byte) [][]byte {
	logs = bytes.ReplaceAll(logs, []byte("\r"), []byte(""))
	return bytes.Split(logs, []byte("\n"))
}
