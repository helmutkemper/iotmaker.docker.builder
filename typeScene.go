package iotmakerdockerbuilder

import "sync"

type scene struct {
	StopedContainers    int
	PausedContainers    int
	MaxStopedContainers int
	MaxPausedContainers int
}

type Theater struct {
	s sync.Mutex
	m map[string]scene
}

func (e *Theater) Init() {
	e.s.Lock()
	defer e.s.Unlock()

	e.m = make(map[string]scene)
}

func (e *Theater) ConfigScene(sceneName string, maxStopedContainers, maxPausedContainers int) {
	e.s.Lock()
	defer e.s.Unlock()

	e.m[sceneName] = scene{
		StopedContainers:    0,
		PausedContainers:    0,
		MaxStopedContainers: maxStopedContainers,
		MaxPausedContainers: maxPausedContainers,
	}
}

func (e *Theater) SetContainerUnPaused(sceneName string) {
	if sceneName == "" {
		return
	}

	e.s.Lock()
	defer e.s.Unlock()

	sc := e.m[sceneName]
	sc.PausedContainers = sc.PausedContainers - 1
	e.m[sceneName] = sc
}

func (e *Theater) SetContainerPaused(sceneName string) {
	if sceneName == "" {
		return
	}

	e.s.Lock()
	defer e.s.Unlock()

	sc := e.m[sceneName]
	sc.PausedContainers = sc.PausedContainers + 1
	e.m[sceneName] = sc
}

func (e *Theater) SetContainerStopped(sceneName string) {
	if sceneName == "" {
		return
	}

	e.s.Lock()
	defer e.s.Unlock()

	sc := e.m[sceneName]
	sc.StopedContainers = sc.StopedContainers + 1
	e.m[sceneName] = sc
}

func (e *Theater) SetContainerUnStopped(sceneName string) {
	if sceneName == "" {
		return
	}

	e.s.Lock()
	defer e.s.Unlock()

	sc := e.m[sceneName]
	sc.StopedContainers = sc.StopedContainers - 1
	e.m[sceneName] = sc
}

var theater = Theater{}

func init() {
	theater.Init()
}
