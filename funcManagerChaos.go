package iotmakerdockerbuilder

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
	"log"
	"strconv"
	"time"
)

func (e *ContainerBuilder) managerChaos() {
	var err error
	var logs []byte
	var lineList [][]byte
	var line []byte
	var found bool
	var timeToNextEvent time.Duration
	var probality float64
	var lineNumber int
	var event Event

	var inspect iotmakerdocker.ContainerInspect

	probality = e.getProbalityNumber()

	inspect, err = e.ContainerInspect()
	if err != nil {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
		event.Error = true
		e.chaos.event <- event
		return
	}

	if e.verifyStatusError(inspect) == true {
		return
	}

	logs, err = e.GetContainerLog()
	if err != nil {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
		event.Error = true
		e.chaos.event <- event
		return
	}

	lineList = e.logsCleaner(logs)

	err = e.writeContainerLogToFile(e.chaos.logPath, lineList)
	if err != nil {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
		event.Error = true
		e.chaos.event <- event
		return
	}

	line, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterFail)
	if found == true {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = string(line)
		event.Fail = true
		e.chaos.event <- event
	}

	line, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterSuccess)
	if found == true {
		_, lineNumber = e.traceCodeLine()
		event.clear()
		event.ContainerName = e.GetContainerName()
		event.Message = string(line)
		event.Done = true
		e.chaos.event <- event
	}

	if e.chaos.enableChaos == false {
		return
	}

	if e.chaos.chaosStarted == false {

		timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToStartChaos, e.chaos.minimumTimeToStartChaos)

		if e.chaos.filterToStart != nil && e.chaos.minimumTimeToStartChaos > 0 {

			_, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterToStart)
			if found == true {
				if e.chaos.serviceStartedAt.Add(timeToNextEvent).Before(time.Now()) == true {
					e.chaos.chaosStarted = true
				}
			}

		} else if e.chaos.filterToStart != nil {

			_, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterToStart)
			if found == true {
				e.chaos.chaosStarted = true
			}

		} else if e.chaos.serviceStartedAt.Add(timeToNextEvent).Before(time.Now()) == true {
			e.chaos.chaosStarted = true
		}

		if e.chaos.chaosStarted == true {
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToStartChaos, e.chaos.minimumTimeToStartChaos)
		} else {
			return
		}

	}

	if e.chaos.chaosCanRestartContainer == false {

		timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeBeforeRestart, e.chaos.minimumTimeBeforeRestart)

		if e.chaos.filterRestart != nil && e.chaos.minimumTimeBeforeRestart > 0 {

			_, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterRestart)
			if found == true {
				if e.chaos.serviceStartedAt.Add(timeToNextEvent).Before(time.Now()) == true {
					e.chaos.chaosCanRestartContainer = true
				}
			}

		} else if e.chaos.filterRestart != nil {

			_, found = e.logsSearchAndReplaceIntoText(lineList, e.chaos.filterRestart)
			if found == true {
				e.chaos.chaosCanRestartContainer = true
			}

		} else if e.chaos.serviceStartedAt.Add(timeToNextEvent).Before(time.Now()) == true {
			e.chaos.chaosCanRestartContainer = true
		}

	}

	//if e.chaos.containerStarted == false {
	//	return
	//}

	//var restartEnable = time.Now().After(e.chaos.minimumTimeBeforeRestart) == true || time.Now().Equal(e.chaos.minimumTimeBeforeRestart) == true

	if time.Now().After(e.chaos.eventNext) == true || time.Now().Equal(e.chaos.eventNext) == true {

		if e.chaos.containerPaused == true {

			log.Printf("%v: unpause()", e.containerName)
			e.chaos.containerPaused = false
			err = e.ContainerUnpause()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToPause, e.chaos.minimumTimeToPause)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		} else if e.chaos.containerStopped == true {

			log.Printf("%v: start()", e.containerName)
			e.chaos.containerStopped = false
			err = e.ContainerStart()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToPause, e.chaos.minimumTimeToPause)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		} else if e.chaos.chaosCanRestartContainer == true && e.chaos.restartProbability != 0.0 && e.chaos.restartProbability >= probality && e.chaos.restartLimit > 0 {

			log.Printf("%v: stop()", e.containerName)
			e.chaos.containerStopped = true
			err = e.ContainerStop()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			e.chaos.restartLimit -= 1
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToRestart, e.chaos.minimumTimeToRestart)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		} else {

			log.Printf("%v: pause()", e.containerName)
			e.chaos.containerPaused = true
			err = e.ContainerPause()
			if err != nil {
				_, lineNumber = e.traceCodeLine()
				event.clear()
				event.ContainerName = e.GetContainerName()
				event.Message = "[" + strconv.Itoa(lineNumber) + "]: " + err.Error()
				event.Error = true
				return
			}
			timeToNextEvent = e.selectBetweenMaxAndMin(e.chaos.maximumTimeToUnpause, e.chaos.minimumTimeToUnpause)
			e.chaos.eventNext = time.Now().Add(timeToNextEvent)

		}
	}
}
