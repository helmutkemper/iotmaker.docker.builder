package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
	"strconv"
)

func (e *ContainerBuilder) writeBlkioIoServiceTimeRecursive(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	if makeLabel == true && e.logFlags&KBlkioIoServiceTimeRecursive == KBlkioIoServiceTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServiceTimeRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Service TimeRecursive.\t"))
			if err != nil {
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoServiceTimeRecursive == KBlkioIoServiceTimeRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServiceTimeRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceTimeRecursive[i].Major, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceTimeRecursive[i].Minor, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(stats.BlkioStats.IoServiceTimeRecursive[i].Op))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServiceTimeRecursive[i].Value, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}
		}
	}

	return
}
