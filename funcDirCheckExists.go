package iotmakerdockerbuilder

import "os"

func DirCheckExists(path string) (exists bool) {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) == false && info.IsDir() == true {
		return true
	}
	return false
}
