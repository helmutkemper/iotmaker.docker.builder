package iotmakerdockerbuilder

import (
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"log"
	"os"
	"strconv"
)

func (e *ContainerBuilder) writeBlkioIoMergedRecursive(file *os.File, stats *types.Stats, makeLabel bool) (err error) {
	if makeLabel == true && e.logFlags&KBlkioIoMergedRecursive == KBlkioIoMergedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoMergedRecursive); i += 1 {
			_, err = file.Write([]byte("BlkioStats stores All IO service stats for data read and write. Io Merged Recursive.\t"))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}
		}
	} else if e.logFlags&KBlkioIoMergedRecursive == KBlkioIoMergedRecursive {
		for i := 0; i != len(stats.BlkioStats.IoMergedRecursive); i += 1 {
			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoMergedRecursive[i].Major, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoMergedRecursive[i].Minor, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(stats.BlkioStats.IoMergedRecursive[i].Op))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}

			_, err = file.Write([]byte(strconv.FormatUint(stats.BlkioStats.IoMergedRecursive[i].Value, 10)))
			if err != nil {
				log.Printf("writeContainerLogToFile().error: %v", err.Error())
				util.TraceToLog()
				return
			}
		}
	}

	return
}
