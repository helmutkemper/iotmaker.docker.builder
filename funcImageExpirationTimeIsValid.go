package iotmakerdockerbuilder

import "time"

func (e *ContainerBuilder) imageExpirationTimeIsValid() (valid bool) {
	if e.imageExpirationTime == 0 {
		return
	}

	var err error
	_, err = e.ImageInspect()
	if err != nil {
		return
	}

	return e.GetImageCreated().Add(e.GetImageExpirationTime()).After(time.Now())
}
