package iotmakerdockerbuilder

import (
	"bytes"
	"github.com/helmutkemper/util"
	"io/fs"
	"io/ioutil"
	"log"
	"strconv"
)

func (e *ContainerBuilder) logsCleaner(logs []byte) [][]byte {

	size := len(logs)

	// faz o log só lê a parte mais recente do mesmo
	logs = logs[e.logsLastSize:]
	e.logsLastSize = size

	// todo: apagar - início
	var dirList []fs.FileInfo
	var err error
	dirList, err = ioutil.ReadDir("./debug_log/")
	if err != nil {
		log.Printf("ioutil.ReadDir().error: %v", err.Error())
		util.TraceToLog()
	}
	var totalOfFiles = strconv.Itoa(len(dirList))
	err = ioutil.WriteFile("./debug_log/"+"log."+totalOfFiles+".log", logs, fs.ModePerm)
	if err != nil {
		log.Printf("ioutil.WriteFile().error: %v", err.Error())
		util.TraceToLog()
	}
	// todo: apagar - fim

	logs = bytes.ReplaceAll(logs, []byte("\r"), []byte(""))
	return bytes.Split(logs, []byte("\n"))
}
