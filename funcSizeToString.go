package iotmakerdockerbuilder

import "strconv"

func (e *ContainerBuilder) SizeToString(value int64) string {

	if value == -1 {
		return "0 B"
	}

	if value > KGigaByte {
		return strconv.FormatFloat(float64(value)/KGigaByte, 'f', 1, 64) + " GB"
	}

	if value > KMegaByte {
		return strconv.FormatFloat(float64(value)/KMegaByte, 'f', 1, 64) + " MB"
	}

	if value > KKiloByte {
		return strconv.FormatFloat(float64(value)/KKiloByte, 'f', 1, 64) + " KB"
	}

	return strconv.FormatInt(value, 10) + " B"
}
