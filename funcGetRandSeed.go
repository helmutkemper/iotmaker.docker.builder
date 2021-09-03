package iotmakerdockerbuilder

import (
	"math/rand"
	"time"
)

func (e *ContainerBuilder) getRandSeed() (seed *rand.Rand) {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}
