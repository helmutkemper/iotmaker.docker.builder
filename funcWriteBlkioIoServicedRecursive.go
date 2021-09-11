package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
	"strconv"
)

func (e *ContainerBuilder) writeBlkioIoServicedRecursive(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	if makeLabel == true && e.logFlags&KBlkioIoServicedRecursive == KBlkioIoServicedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServicedRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Serviced Recursive.\t"))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoServicedRecursive == KBlkioIoServicedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoServicedRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServicedRecursive[i].Major, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServicedRecursive[i].Minor, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(stats.BlkioStats.IoServicedRecursive[i].Op))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoServicedRecursive[i].Value, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}
		}
	}

	return
}
